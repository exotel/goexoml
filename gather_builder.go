package goexoml

//SetAction sets Action
func (__gather__ *Gather) SetAction(action string) *Gather {
	__gather__.Action = action
	return __gather__
}

//SetMethod sets Method
func (__gather__ *Gather) SetMethod(method string) *Gather {
	__gather__.Method = method
	return __gather__
}

//SetTimeout sets Timeout
func (__gather__ *Gather) SetTimeout(timeout int) *Gather {
	__gather__.Timeout = timeout
	return __gather__
}

//SetFinishOnKey sets FinishOnKey
func (__gather__ *Gather) SetFinishOnKey(finishonkey string) *Gather {
	__gather__.FinishOnKey = finishonkey
	return __gather__
}

//SetNumDigits sets NumDigits
func (__gather__ *Gather) SetNumDigits(numdigits int) *Gather {
	__gather__.NumDigits = numdigits
	return __gather__
}

//SetSay sets Say
func (__gather__ *Gather) SetSay(say Say) *Gather {
	__gather__.Say = say
	return __gather__
}

//SetPlay sets Play
func (__gather__ *Gather) SetPlay(play Play) *Gather {
	__gather__.Play = play
	return __gather__
}

//NewGather return a new Gather pointer
func NewGather() *Gather {
	return new(Gather)
}

//AddGather appends the verb to response
func (r *Response) AddGather(gather Gather) *Response {
	r.Response = append(r.Response, gather)
	return r
}
