package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

// implement Handler ServeHTTP
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello foo:)")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Golnag web:)")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello bar:)")
	})

	// add handler instance
	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}
