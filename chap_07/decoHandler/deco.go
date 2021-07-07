package decohandler

import "net/http"

// DecoratorFunc Decorator
type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

// DecoHandler Decorator
type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

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
