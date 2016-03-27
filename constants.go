package goexoml

//CallStatus has the Constants for CallStatus String
var CallStatus = struct {
	EXOQueued     string
	EXORinging    string
	EXOInProgress string
	EXOCompleted  string
	EXOBusy       string
	EXOFailed     string
	EXONoAnswer   string
	EXOCanceled   string
}{
	EXOQueued:     "queued",
	EXORinging:    "ringing",
	EXOInProgress: "in-progress",
	EXOCompleted:  "completed",
	EXOBusy:       "busy",
	EXOFailed:     "failed",
	EXONoAnswer:   "no-answer",
	EXOCanceled:   "canceled",
}
