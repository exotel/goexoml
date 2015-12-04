package goexoml

//SetAction sets Action
func (__record__ *Record) SetAction(action string) *Record {
	__record__.Action = action
	return __record__
}

//SetMethod sets Method
func (__record__ *Record) SetMethod(method string) *Record {
	__record__.Method = method
	return __record__
}

//SetTimeout sets Timeout
func (__record__ *Record) SetTimeout(timeout int) *Record {
	__record__.Timeout = timeout
	return __record__
}

//SetFinishOnKey sets FinishOnKey
func (__record__ *Record) SetFinishOnKey(finishonkey string) *Record {
	__record__.FinishOnKey = finishonkey
	return __record__
}

//SetMaxLength sets MaxLength
func (__record__ *Record) SetMaxLength(maxlength int) *Record {
	__record__.MaxLength = maxlength
	return __record__
}

//SetTranscribe sets Transcribe
func (__record__ *Record) SetTranscribe(transcribe bool) *Record {
	__record__.Transcribe = transcribe
	return __record__
}

//SetTranscribeCallback sets TranscribeCallback
func (__record__ *Record) SetTranscribeCallback(transcribecallback string) *Record {
	__record__.TranscribeCallback = transcribecallback
	return __record__
}

//SetPlayBeep sets PlayBeep
func (__record__ *Record) SetPlayBeep(playbeep bool) *Record {
	__record__.PlayBeep = playbeep
	return __record__
}

//SetTrim sets Trim
func (__record__ *Record) SetTrim(trim string) *Record {
	__record__.Trim = trim
	return __record__
}

//NewRecord return a new Record pointer
func NewRecord() *Record {
	return new(Record)
}

//AddRecord appends the verb to response
func (r *Response) AddRecord(record Record) *Response {
	r.Response = append(r.Response, record)
	return r
}
