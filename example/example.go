package main

import (
	"fmt"

	"github.com/exotel/goexoml"
)

//Dial - creates a dial verb inside response
func Dial(number string, action string, callerID string) (string, error) {
	resp := goexoml.NewResponse()
	dial := goexoml.NewDial().
		SetAction(action).
		SetCallerID(callerID)
	num := goexoml.NewNumber().SetNoun(number)
	dial.SetNumber(*num)
	resp.AddDial(*dial)
	resp.AddRecord(goexoml.Record{})
	return resp.String(), nil
}

func main() {
	dial, err := Dial("sasa", "sasasa", "sasass")
	if err != nil {
		fmt.Print("Error occured :", err.Error())
	}
	fmt.Println(dial)
	return
}
