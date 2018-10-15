package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

// Struct defines the field set array for the struct
type Struct struct {
	Name   string
	Fields []Field
}

// Field defines the Fields pf
type Field struct {
	FieldName string
	FieldType string
}

var (
	fileName   = flag.String("f", "", "The source file path")
	structName = flag.String("t", "", "The name of the struct for which the files has to be generated")
)

func getTemplate() (temp *template.Template, err error) {

	var handler = `
	// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 // Autogenerated by buildergenerator

	package goexoml
	import (
		"errors"
	)
	var _ = errors.New("_")
  {{$name := .Name}}
	{{range $index, $field := .Fields}}
	// Set{{$field.FieldName}} sets {{$field.FieldName}} for {{$name}} struct instance
	func ({{$name | receiver}} *{{$name}}) Set{{$field.FieldName}}({{$field.FieldName | lower}} {{$field.FieldType}}) *{{$name}}{
		{{$name | receiver}}.{{$field.FieldName}} = {{$field.FieldName | lower}}
		return {{$name | receiver}}
	}
	{{end}}

  // Setter returns setter function for the field given
	func ({{$name | receiver}} *{{$name}}) Setter(field string) (setter func(interface{}) (*{{$name}}, error)) {
		switch field { {{range $index, $field := .Fields}}
	case "{{$field.FieldName}}":
		setter = func({{$field.FieldName}}Field interface{}) (*{{$name}}, error) {
			if {{$field.FieldName}}Value, ok := {{$field.FieldName}}Field.({{$field.FieldType}}); ok {
				return {{$name | receiver}}.Set{{$field.FieldName}}({{$field.FieldName}}Value), nil
			}
			return nil, errors.New("Invalid type Expected {{$field.FieldType}}")
		}	{{end}}
		}
		return
	}
  // New{{.Name}} return a new {{.Name}} pointer
	func New{{.Name}}() *{{.Name}} {
		return new({{.Name}})
	}
	// I{{.Name}} The interface that satisfies all the methods for this struct
	// I{{.Name}} asserts implementation of setters for all the fields of {{.Name}}
  type I{{.Name}}	interface {
	    {{range $index, $field := .Fields}}Set{{$field.FieldName}}({{$field.FieldName | lower}} {{$field.FieldType}}) *{{$name}};{{end}}
			Setter(string) (func(interface{}) (*{{$name}}, error))
  }
	// Add{{.Name}} appends the verb to response
	func (r *Response) Add{{.Name}}({{.Name|lower}} I{{.Name}}) *Response {
		r.Response = append(r.Response,{{.Name|lower}})
		return r
	}
	`
	// Create a new template and parse the letter into it.
	t := template.
		Must(template.
		New("builder").
		Funcs(template.FuncMap{"lower": strings.ToLower, "receiver": func(s string) string { return "__" + strings.ToLower(s) + "__" }}).
		Parse(handler))
	temp = t
	return

}

func generateCode(t *template.Template, str Struct) (code []byte, err error) {
	var bW bytes.Buffer
	w := bufio.NewWriter(&bW)
	err = t.Execute(w, str)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}
	fmt.Println(string(bW.Bytes()))
	//format the bytes from bufffer
	code, err = format.Source(bW.Bytes())
	return
}

func generateBuilder(str Struct) (err error) {
	var t *template.Template
	t, err = getTemplate()

	if err != nil {
		return
	}
	//create the template into a bytes Buffer
	var code []byte
	code, err = generateCode(t, str)
	if err != nil {
		return
	}

	//write to a file
	var f *os.File
	f, err = os.Create(strings.ToLower(str.Name) + "_builder.go")
	if err != nil {
		return
	}
	_, err = bytes.NewReader(code).WriteTo(f)
	if err != nil {
		return
	}
	return
}

func structFields(structName string, f *ast.File, src string) (fs []Field) {
	fields := f.Scope.Lookup(structName).Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields
	for _, field := range fields.List {
		if field.Names[0].String() != "XMLName" {
			fs = append(fs, Field{FieldName: field.Names[0].String(), FieldType: src[field.Type.Pos()-1 : field.Type.End()]})
		}
	}
	return
}

func getStructInfo(filename string, structName string) (str Struct, err error) {
	var bf bytes.Buffer
	fset := token.NewFileSet()
	var file *os.File
	file, err = os.Open(filename)
	if err != nil {
		return
	}
	bf.ReadFrom(file)
	b := bf.Bytes()
	var f *ast.File
	f, err = parser.ParseFile(fset, "", b, 0)
	if err != nil {
		return
	}
	str.Name = structName
	str.Fields = structFields(structName, f, string(b))
	return
}

func createStruct() (str Struct, err error) {
	flag.Parse()
	str, err = getStructInfo(*fileName, *structName)
	return
}

func main() {
	str, err := createStruct()
	if err != nil {
		log.Fatal("Error happened creating the resource information", err.Error())
	}
	if err = generateBuilder(str); err != nil {
		log.Fatalln("Error happened generating the builder ", err.Error())
	}
}
