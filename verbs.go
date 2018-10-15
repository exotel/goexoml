package goexoml

import "encoding/xml"

// Dial struct for the verb Dial which allows a number as plain test
//go:generate buildergenerator -t Dial -f $GOFILE
type Dial struct {
	XMLName      xml.Name `xml:"Dial"`
	Action       string   `xml:"action,attr,omitempty"`
	Method       string   `xml:"method,attr,omitempty"`
	Timeout      int      `xml:"timeout,attr,omitempty"`
	HangupOnStar bool     `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    int      `xml:"timeLimit,attr,omitempty"`
	CallerID     string   `xml:"callerId,attr,omitempty"`
	Record       bool     `xml:"record,attr,omitempty"`
	PlainNumber  string   `xml:",innerxml"`
	Number       *Number  `xml:",innerxml"`
}

// Hangup struct for the verb Hangup
//go:generate buildergenerator -t Hangup -f $GOFILE
type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

// Play struct for the verb Play
//go:generate buildergenerator -t Play -f $GOFILE
type Play struct {
	XMLName xml.Name `xml:"Play"`
	Loop    int      `xml:"loop,attr,omitempty"`
	Digits  int      `xml:"digits,attr,omitempty"`
	URL     string   `xml:",innerxml"`
}

// Record struct for the verb Record
//go:generate buildergenerator -t Record -f $GOFILE
type Record struct {
	XMLName            xml.Name `xml:"Record"`
	Action             string   `xml:"action,attr,omitempty"`
	Method             string   `xml:"method,attr,omitempty"`
	Timeout            int      `xml:"timeout,attr,omitempty"`
	FinishOnKey        string   `xml:"finishOnKey,attr,omitempty"`
	MaxLength          int      `xml:"maxLength,attr,omitempty"`
	Transcribe         bool     `xml:"transcribe,attr,omitempty"`
	TranscribeCallback string   `xml:"transcribeCallback,attr,omitempty"`
	PlayBeep           bool     `xml:"playBeep,attr,omitempty"`
	Trim               string   `xml:"trim,attr,omitempty"`
}

// Redirect struct for the verb Redirect
//go:generate buildergenerator -t Redirect -f $GOFILE
type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Method  string   `xml:"method,attr,omitempty"`
	URL     string   `xml:",innerxml"`
}

// Say struct for the verb Say
//go:generate buildergenerator -t Say -f $GOFILE
type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`
	Text     string   `xml:",innerxml"`
}

// Gather struct for the verb Gather
//go:generate buildergenerator -t Gather -f $GOFILE
type Gather struct {
	XMLName     xml.Name `xml:"Gather"`
	Action      string   `xml:"action,attr,omitempty"`
	Method      string   `xml:"method,attr,omitempty"`
	Timeout     int      `xml:"timeout,attr,omitempty"`
	FinishOnKey string   `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int      `xml:"numDigits,attr,omitempty"`
	Say         *Say     `xml:",innerxml"`
	Play        *Play    `xml:",innerxml"`
}

// Number struct for the noun Number
//go:generate buildergenerator -t Number -f $GOFILE
type Number struct {
	XMLName              xml.Name `xml:"Number"`
	SendDigits           string   `xml:"sendDigits,attr,omitempty"`
	URL                  string   `xml:"URL,attr,omitempty"`
	Method               string   `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string   `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string   `xml:"StatusCallback,attr,omitempty"`
	StatusCallbackMethod string   `xml:"StatusCallbackMethod,attr,omitempty"`
	Noun                 string   `xml:",innerxml"`
}
