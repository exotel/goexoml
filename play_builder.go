package goexoml

//SetLoop sets Loop
func (__play__ *Play) SetLoop(loop int) *Play {
	__play__.Loop = loop
	return __play__
}

//SetDigits sets Digits
func (__play__ *Play) SetDigits(digits int) *Play {
	__play__.Digits = digits
	return __play__
}

//SetURL sets URL
func (__play__ *Play) SetURL(url string) *Play {
	__play__.URL = url
	return __play__
}

//NewPlay return a new Play pointer
func NewPlay() *Play {
	return new(Play)
}

//AddPlay appends the verb to response
func (r *Response) AddPlay(play Play) *Response {
	r.Response = append(r.Response, play)
	return r
}
