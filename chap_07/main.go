package main

import (
	app "golang_web/chap_07/app"
	"net/http"
)

func main() {
	mux := app.NewHandler()
	http.ListenAndServe(":3000", mux)
}
