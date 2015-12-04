package goexoml

//SetAction sets Action
func (__dial__ *Dial) SetAction(action string) *Dial {
	__dial__.Action = action
	return __dial__
}

//SetMethod sets Method
func (__dial__ *Dial) SetMethod(method string) *Dial {
	__dial__.Method = method
	return __dial__
}

//SetTimeout sets Timeout
func (__dial__ *Dial) SetTimeout(timeout int) *Dial {
	__dial__.Timeout = timeout
	return __dial__
}

//SetHangupOnStar sets HangupOnStar
func (__dial__ *Dial) SetHangupOnStar(hanguponstar bool) *Dial {
	__dial__.HangupOnStar = hanguponstar
	return __dial__
}

//SetTimeLimit sets TimeLimit
func (__dial__ *Dial) SetTimeLimit(timelimit int) *Dial {
	__dial__.TimeLimit = timelimit
	return __dial__
}

//SetCallerID sets CallerID
func (__dial__ *Dial) SetCallerID(callerid string) *Dial {
	__dial__.CallerID = callerid
	return __dial__
}

//SetRecord sets Record
func (__dial__ *Dial) SetRecord(record bool) *Dial {
	__dial__.Record = record
	return __dial__
}

//SetPlainNumber sets PlainNumber
func (__dial__ *Dial) SetPlainNumber(plainnumber string) *Dial {
	__dial__.PlainNumber = plainnumber
	return __dial__
}

//SetNumber sets Number
func (__dial__ *Dial) SetNumber(number Number) *Dial {
	__dial__.Number = number
	return __dial__
}

//NewDial return a new Dial pointer
func NewDial() *Dial {
	return new(Dial)
}

//AddDial appends the verb to response
func (r *Response) AddDial(dial Dial) *Response {
	r.Response = append(r.Response, dial)
	return r
}
