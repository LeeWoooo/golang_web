package main

import (
	"golang_web/chap_05/app"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", app.NewHandler())
}
