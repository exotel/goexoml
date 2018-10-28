package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestStructFields(t *testing.T) {
	// test structFields()
	src := `package foo

type Foo struct {
  	XMLName xml.Name` + " `" + `xml:"Play"` + "`\n" +
		`Loop    int` + " `" + `xml:"loop,attr,omitempty"` + "`\n" +
		`Digits  int` + " `" + `xml:"digits,attr,omitempty"` + "`\n" +
		`}`

	f, err := parser.ParseFile(token.NewFileSet(), "", src, 0)
	if err != nil {
		t.Fatal("parse error", err)
	}

	fields := structFields("Foo", f, src)

	// 'loop' field
	if have, want := fields[0].FieldName, "Loop"; have != want {
		t.Errorf("have = %v want = %s", have, want)
	}

	if have, want := fields[0].FieldType, "int"; have != want {
		t.Errorf("have = %v want = %s", have, want)
	}

	// 'digits' field
	if have, want := fields[1].FieldName, "Digits"; have != want {
		t.Errorf("have = %v want = %s", have, want)
	}

	if have, want := fields[1].FieldType, "int"; have != want {
		t.Errorf("have = %v want = %s", have, want)
	}

	// 'XMLName should not be included
	for _, f := range fields {
		if f.FieldName == "XMLName" {
			t.Error("unexpected 'XMLName' in the fields")
		}
	}
}

func TestGenerateCode(t *testing.T) {
	strct := Struct{
		Name: "Play",
		Fields: []Field{
			Field{
				FieldName: "Loop",
				FieldType: "int",
			},
			Field{
				FieldName: "Digits",
				FieldType: "int",
			},
		},
	}

	tpl, err := getTemplate()
	if err != nil {
		t.Error(err)
	}

	srcbuf, err := generateCode(tpl, strct)
	if err != nil {
		t.Error(err)
	}

	// 1. test is if generated code is parsable
	af, err := parser.ParseFile(token.NewFileSet(), "", srcbuf, 0)
	if err != nil {
		t.Fatal("generated code is not parsable", err)
	}

	// 2. check for obvious generated function decls
	expectedDecls := make(map[string]bool)
	expectedDecls["SetLoop"] = false
	expectedDecls["SetDigits"] = false
	expectedDecls["Setter"] = false
	expectedDecls["NewPlay"] = false
	expectedDecls["AddPlay"] = false

	for _, decl := range af.Decls {
		if d, ok := decl.(*ast.FuncDecl); ok {
			expectedDecls[d.Name.String()] = true
		}
	}

	for k, v := range expectedDecls {
		if v == false {
			t.Errorf("expected decl '%s', not found", k)
		}
	}
}
