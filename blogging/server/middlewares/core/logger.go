package core

import (
	"log"
	"net/http"
	"time"
)

var Logger = Middleware{
	Operation: Include,
	Func: func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Println(r.Method, r.URL.Path, time.Since(start))
		})
	},
	Urls: map[string]struct{}{},
}
