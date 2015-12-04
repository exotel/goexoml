package goexoml

//NewHangup return a new Hangup pointer
func NewHangup() *Hangup {
	return new(Hangup)
}

//AddHangup appends the verb to response
func (r *Response) AddHangup(hangup Hangup) *Response {
	r.Response = append(r.Response, hangup)
	return r
}
