package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "index")
}

// NewHandler make Handler instance
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandle)
	h := NewDecoHandler(mux, loggerDecorator)
	return h
}

/*
DecoratorFunc Decorator
response,request,handler(다음으로 실행될 handler)를 가지고있다.
이 함수를 구현한 함수는 request와 response를 가지고 다음으로 실행될 handler의 ServeHttp를 실행시킨다.
ServeHttp를 실행시키기 전 앞 뒤로 Decorator를 추가할 수 있다.
*/
type DecoratorFunc func(w http.ResponseWriter, r *http.Request, h http.Handler)

/*
DecoHandler Decorator
DecoratorHandler는 http Handler와 fn을 가지고 있다.
DecoratorHandler를 Handler로 등록한 webServer에 요청이 들어오면 ServeHTTP를 실행한다.
*/
type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

/*
http.Handler implement
*/
func (d *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.fn(w, r, d.h)
}

// NewDecoHandler Create DecoHandler instance
func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  h,
	}
}

func loggerDecorator(w http.ResponseWriter, r *http.Request, h http.Handler) {
	now := time.Now()
	log.Println("Start")
	h.ServeHTTP(w, r)
	log.Println("End", time.Since(now).Milliseconds())
}
