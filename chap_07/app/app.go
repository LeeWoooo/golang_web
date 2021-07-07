package app

import (
	"fmt"
	decohandler "golang_web/chap_07/decoHandler"
	"log"
	"net/http"
	"time"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] End", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] End", time.Since(start).Milliseconds())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Go")
}

// NewHandler create mux instance
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	h := decohandler.NewDecoHandler(mux, logger)
	h = decohandler.NewDecoHandler(mux, logger2)
	//add handler
	mux.HandleFunc("/", indexHandler)

	return h
}
