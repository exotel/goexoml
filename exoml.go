package goexoml

import (
	"encoding/xml"
	"fmt"
)

// Action appends action verb structs to response.
func (r *Response) Action(structs ...interface{}) error {
	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("non valid verb: '%T'", s)
		case *Hangup, Hangup, Play, *Play, Record, *Record, *Say, *Redirect, Redirect, *Gather, Gather, *Dial, Dial:
			r.Response = append(r.Response, s)
		}
	}
	return nil
}

// String returns a formatted xml response
func (r Response) String() string {
	output, err := xml.MarshalIndent(r, "  ", "    ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return xml.Header + string(output) + "\n"
}
