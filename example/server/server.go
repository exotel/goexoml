package main

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/exotel/goexoml"
	"github.com/go-chi/chi"
	mw "github.com/go-chi/chi/middleware"
)

//Dial returns an exoml given number to call a number
func Dial(w http.ResponseWriter, r *http.Request) {
	var number string
	number = chi.URLParam(r, "number")
	if number == "" {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("PROVIDE NUMBER TO BE CALLED"))
	}

	resp := goexoml.NewResponse()
	say := goexoml.NewSay().SetText("You can't handle the truth")
	dial := goexoml.NewDial().SetPlainNumber(number)
	resp.AddSay(say).AddDial(dial).AddHangup(goexoml.NewHangup())

	xml, err := xml.Marshal(resp)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("content-type", "application/xml;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(xml)
}

func main() {
	r := chi.NewRouter()
	r.Use(mw.Logger)
	r.Use(mw.Recoverer)

	r.Get("/dial/{number}", Dial)

	log.Fatal(http.ListenAndServe(":1323", r))
}
