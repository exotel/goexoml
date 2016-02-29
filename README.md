### goexoml [![GoDoc](https://godoc.org/gopkg.in/exotel/goexoml.v1?status.svg)](https://godoc.org/gopkg.in/exotel/goexoml.v1)

[![Join the chat at https://gitter.im/exotel/goexoml](https://badges.gitter.im/exotel/goexoml.svg)](https://gitter.im/exotel/goexoml?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
The `ExoML` library for golang

####What is it?
  goexoml is the official ExoML library written in golang. <br>**ExoML** (ExotelMarkupLanguage) enables one to add logic to calls.When someone makes a call  to an *exophone* ,Exotel will look up the URL associated with *exophone* and make a request to that URL.
	The URL can be configured to respond with ExoMLs Responses,which exotel interprets and executes


For Example

The following will be interpreted by exotel as a request to make a call to the number +919742033616 and makes call

```
<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Dial>
			<Number>
					+919742033616
			</Number>
	</Dial>
</Response>

```

####currently supported verbs
* Dial  
* Hangup
* Play
* Record
* Redirect
* Say
* Gather


####how to use
Install the latest library using `go get`
```
go get github.com/exotel/goexoml
```

or  from the gopkg repo to make sure the version number is maintained all time
```
go get gopkg.in/exotel/goexoml.v1 //for v1.x.x
```


###sample Code
```
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

	//create the recors vberb and add attributes
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
####What if i don't have go installed ?
This is how you can install go in linux [this is for version 1.5.2 of golang]
```
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

now you have go installed,try checking by checking the version
```
go version
```

Now you have to set the Workspace,say you are setting,~/go as the worksapce
this is how it is done
```
export GOPATH=$HOME/go
```

Add these two  export statements in ~/.bashrc so that its set all the time
Now you are ready to test goexoml,get the library using go get
ie,
```
go get github.com/exotel/goexoml
```


Now you can copy paste the sample code server.go from example/server to a file ,say `path/to/file.go`
The example program uses an external library for spawning http server intsll it as
```
go get github.com/labstack/echo
```

Run the server as

```
go run path/to/file.go
```
By now the exoml library would be running on port 1323

Test it
````
curl http://localhost:1323/dial/+919742033616
```

This would return an exoml
```
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>You can't handle the truth</Say>
    <Dial>+919742033616</Dial>
    <Hangup></Hangup>
</Response
```


####Versioning
#####*major version Changes would mean following *
* Removing or renaming *any* exposed name (function, method, type, etc)
* Adding, removing or renaming methods in an interface
* Adding a parameter to a function, method, or interface
* Changing the type of a parameter or result in a function, method, or interface
* Changing the number of results in a function, method, or interface
* Some user facing struct changes

#####*expect these changes even without major version being same *
* Adding exposed names (function, method, type, etc)
* Renaming a parameter or result of a function, method, or interface*
* Some struct changes which does not directly affect the functionality of the library [*much*] but may/maynot add new features to the existing library





####contributions
sarath@exotel.in
