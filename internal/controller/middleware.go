package controller

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func wrapWithMiddleware(router http.Handler) http.Handler {
	return cors.Default().Handler(withLogging(router))
}

func withLogging(h http.Handler) http.Handler {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(rw, r) // serve the original request

		log.Printf("%s [%s]", method, uri)
	}

	return http.HandlerFunc(logFn)
}
