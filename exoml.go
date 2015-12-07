//Package goexoml The ExoML library for golang
//goexoml is the official ExoML library written in golang. ExoML(ExotelMarkupLanguage) enables one to add logic to calls.
//
//When someone makes a call  to an *exophone* ,Exotel will look up the URL associated with *exophone* and make a request to that URL.
//
//sample Code
//	 package main
//
//	 import (
//	 	"fmt"
//
//	 	"github.com/exotel/goexoml"
//	 )
//
//	 //Example - creates a sample response as follows
//	 // <?xml version="1.0" encoding="UTF-8"?>
//	 // <Response>
//	 //     <Say>
//	 //         Please leave a message at the beep.
//	 //         Press the star key when finished.
//	 //     </Say>
//	 //     <Record
//	 //         action="http://example.exotel.in/handleRecording"
//	 //         method="GET"
//	 //         maxLength="20"
//	 //         finishOnKey="*"
//	 //         />
//	 //     <Say>I did not receive a recording</Say>
//	 // </Response>
//	 func Example() (string, error) {
//	 	//create a new response object
//	 	resp := goexoml.NewResponse()
//
//	 	//create a new say verb and add attributes and values
//	 	say1 := goexoml.NewSay().SetText("Please leave a message at the beep.\n         Press the star key when finished.")
//
//	 	//create a new say verb and add attributes and values
//	 	say2 := goexoml.NewSay().SetText("I did not receive a recording.")
//
//	 	//create the recors vberb and add attributes
//	 	rec := goexoml.NewRecord().SetAction("http://example.exotel.in/handleRecording").SetMethod("GET").SetMaxLength(20)
//
//	 	//Add the Action verbs to the response object in expected order
//	 	err := resp.Action(say1, rec, say2)
//	 	if err != nil {
//	 		fmt.Println(err.Error())
//	 		return "", err
//	 	}
//	 	// OR
//	 	////Add the Action verbs to the response object in expected order
//	 	// resp.AddSay(say1).AddRecord(rec).AddSay(say2)
//	 	return resp.String(), nil
//	 }
//
//	 func main() {
//	 	exoml, err := Example()
//	 	if err != nil {
//	 		fmt.Print("Error occured :", err.Error())
//	 		return
//	 	}
//	 	fmt.Println(exoml)
//	 	return
//	 }
//
package goexoml

import (
	"encoding/xml"
	"fmt"
)

// Action appends action verb structs to response.
//The verbs has to be given in the order in which
//they are expected to be executed
// if there is any invalid verb the function will repond with error but still Response would have all the other
//verbs till the invalid one
func (r *Response) Action(structs ...interface{}) error {
	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("non valid verb: '%T'", s)
		case IHangup, IPlay, IRecord, ISay, IRedirect, IGather, IDial:
			r.Response = append(r.Response, s)
		}
	}
	return nil
}

// String returns a formatted xml response
// String implements the fmt.Stringer and it returns the ExoML Response struct as an XMLMarshalled string
func (r Response) String() string {
	output, err := xml.MarshalIndent(r, "", "   ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}
	return xml.Header + string(output) + "\n"
}
