package main

import (
	"net/http"

	"github.com/exotel/goexoml"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

//Dial returns an exoml given number to call a number
func Dial(c *echo.Context) error {
	var number string
	number = c.Param("number")
	if number == "" {
		return c.String(http.StatusBadRequest, "PROVIDE NUMBER TO BE CALLED")
	}
	resp := goexoml.NewResponse()
	say := goexoml.NewSay().SetText("You can't handle the truth")
	dial := goexoml.NewDial().SetPlainNumber(number) //SetNumber(*(goexoml.NewNumber().SetNoun(number)))
	resp.AddSay(say).AddDial(dial).AddHangup(goexoml.NewHangup())
	return c.XML(http.StatusOK, resp)
}

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Get("/dial/:number", Dial)
	// Start server
	e.Run(":1323")
}
