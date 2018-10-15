### goexoml [![GoDoc](https://godoc.org/gopkg.in/exotel/goexoml.v1?status.svg)](https://godoc.org/gopkg.in/exotel/goexoml.v1) &nbsp;[![Join the chat at https://gitter.im/exotel/goexoml](https://badges.gitter.im/exotel/goexoml.svg)](https://gitter.im/exotel/goexoml?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)


The **`ExoML`** library for golang

#### What is it?
  goexoml is the official ExoML library written in golang. <br>**ExoML** (ExotelMarkupLanguage) enables one to add logic to calls.When someone makes a call  to an *exophone* ,Exotel will look up the URL associated with *exophone* and make a request to that URL.
	The URL can be configured to respond with ExoMLs Responses,which exotel interprets and executes


For Example

The following will be interpreted by exotel as a request to make a call to the number +919742033616 and makes call

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
 <Dial>
   <Number>
     +919742033616
   </Number>
 </Dial>
</Response>

```

#### currently supported verbs
* Dial  
* Hangup
* Play
* Record
* Redirect
* Say
* Gather


#### how to use
Install the latest library using `go get`
```sh
go get github.com/exotel/goexoml
```

or  from the gopkg repo to make sure the version number is maintained all time
```sh
go get gopkg.in/exotel/goexoml.v1 //for v1.x.x
```


### sample Code
```go
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
//         action="http://example.exotel.in/handleRecording"
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

	//create the record verb and add attributes
	rec := goexoml.NewRecord().SetAction("http://example.exotel.in/handleRecording").SetMethod("GET").SetMaxLength(20)

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

```

#### What if i don't have go installed ?
Follow the instructions [here](https://golang.org/doc/install) to install go 

### Example
[example ](https://github.com/exotel/goexoml/tree/master/example) directory contains code demonstrating the library usage

For a quick start you may try the [example/server](https://github.com/exotel/goexoml/tree/master/example/server) 

```sh
# install dependencies
go get github.com/exotel/goexoml/example/server

# run server
go run github.com/exotel/goexoml/example/server
```
By now the exoml server would be running on port 1323

```sh
curl http://localhost:1323/dial/+919742033616
```

This would return an exoml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>You can't handle the truth</Say>
    <Dial>+919742033616</Dial>
    <Hangup></Hangup>
</Response
```


#### Versioning
##### major version Changes would mean following 
* Removing or renaming *any* exposed name (function, method, type, etc)
* Adding, removing or renaming methods in an interface
* Adding a parameter to a function, method, or interface
* Changing the type of a parameter or result in a function, method, or interface
* Changing the number of results in a function, method, or interface
* Some user facing struct changes

##### expect these changes even without major version being same 
* Adding exposed names (function, method, type, etc)
* Renaming a parameter or result of a function, method, or interface*
* Some struct changes which does not directly affect the functionality of the library [*much*] but may/maynot add new features to the existing library



#### contributions

* Sarath S Pillai ([sarathsp06](https://github.com/sarathsp06))
