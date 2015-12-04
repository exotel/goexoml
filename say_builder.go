package goexoml

//SetVoice sets Voice
func (__say__ *Say) SetVoice(voice string) *Say {
	__say__.Voice = voice
	return __say__
}

//SetLanguage sets Language
func (__say__ *Say) SetLanguage(language string) *Say {
	__say__.Language = language
	return __say__
}

//SetLoop sets Loop
func (__say__ *Say) SetLoop(loop int) *Say {
	__say__.Loop = loop
	return __say__
}

//SetText sets Text
func (__say__ *Say) SetText(text string) *Say {
	__say__.Text = text
	return __say__
}

//NewSay return a new Say pointer
func NewSay() *Say {
	return new(Say)
}

//AddSay appends the verb to response
func (r *Response) AddSay(say Say) *Response {
	r.Response = append(r.Response, say)
	return r
}
