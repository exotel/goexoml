package goexoml

import "encoding/xml"

//Response struct for the verb Response
type Response struct {
	XMLName  xml.Name `xml:"Response"`
	Response []interface{}
}

//NewResponse returns a pointer to the reposne structure
func NewResponse() *Response {
	return new(Response)
}
