package core

import (
	"log"
	"matwa/blogger/data"
	"net/http"
	"time"
)

type MiddlewareFunc func(http.Handler) http.Handler

func CreateStack(xs ...MiddlewareFunc) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			h = xs[i](h)
		}
		return h
	}
}

var Logging MiddlewareFunc = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start).Nanoseconds(), "ns")
	})

}

var Authentication MiddlewareFunc = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL != nil {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("token")
		if token == "" {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token is required"))
			return
		}

		for _, user := range data.UsersStore {
			if user.Token == token {

				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid token"))
	})
}
