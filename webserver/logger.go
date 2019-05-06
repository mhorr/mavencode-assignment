package main

import (
	"log"
	"net/http"
	"time"
)

// Logger wraps an http handler and logs it.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			"about to serve",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			"done serving",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
