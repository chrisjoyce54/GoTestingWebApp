package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprint(writer, "This is the home page.")
}
