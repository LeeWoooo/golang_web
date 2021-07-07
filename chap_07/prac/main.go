package main

import (
	"golang_web/chap_07/prac/app"
	"net/http"
)

func main() {
	mux := app.NewHandler()

	http.ListenAndServe(":3000", mux)
}
