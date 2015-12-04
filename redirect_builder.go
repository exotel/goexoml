package goexoml

//SetMethod sets Method
func (__redirect__ *Redirect) SetMethod(method string) *Redirect {
	__redirect__.Method = method
	return __redirect__
}

//SetURL sets URL
func (__redirect__ *Redirect) SetURL(url string) *Redirect {
	__redirect__.URL = url
	return __redirect__
}

//NewRedirect return a new Redirect pointer
func NewRedirect() *Redirect {
	return new(Redirect)
}

//AddRedirect appends the verb to response
func (r *Response) AddRedirect(redirect Redirect) *Response {
	r.Response = append(r.Response, redirect)
	return r
}
