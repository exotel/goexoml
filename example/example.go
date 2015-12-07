package main

import (
	"fmt"

	"github.com/exotel/goexoml"
)

//Example - creates a sample response as follows
// <?xml version="1.0" encoding="UTF-8"?>
// <Response>
//     <Say>
//         Please leave a message at the beep.
//         Press the star key when finished.
//     </Say>
//     <Record
//         action="http://foo.edu/handleRecording.php"
//         method="GET"
//         maxLength="20"
//         finishOnKey="*"
//         />
//     <Say>I did not receive a recording</Say>
// </Response>
func Example() (string, error) {
	//create a new response object
	resp := goexoml.NewResponse()

	//create a new say verb and add attributes and values
	say1 := goexoml.NewSay().SetText("Please leave a message at the beep.\n         Press the star key when finished.")

	//create a new say verb and add attributes and values
	say2 := goexoml.NewSay().SetText("I did not receive a recording.")

	//create the recors vberb and add attributes
	rec := goexoml.NewRecord().SetAction("http://foo.edu/handleRecording.php").SetMethod("GET").SetMaxLength(20)

	//Add the Action verbs to the response object in expected order
	err := resp.Action(say1, rec, say2)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	// OR
	////Add the Action verbs to the response object in expected order
	// resp.AddSay(say1).AddRecord(rec).AddSay(say2)
	return resp.String(), nil
}

func main() {
	exoml, err := Example()
	if err != nil {
		fmt.Print("Error occured :", err.Error())
		return
	}
	fmt.Println(exoml)
	return
}
