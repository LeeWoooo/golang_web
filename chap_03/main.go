package main

import (
	"golang_web/chap_03/myapp"
	"net/http"
)

func main() {
	mux := myapp.NewHandler()
	http.ListenAndServe(":3000", mux)
}
