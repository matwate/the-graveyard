package main

import (
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

type Middleware func(http.Handler) http.Handler

func ChainMiddlewares(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}

}

var logger Middleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %s %s", r.RemoteAddr, r.URL)
		next.ServeHTTP(w, r)
	})
}

var midldewares = ChainMiddlewares(logger)

func (s *Server) Start() error {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))

	})

	server := &http.Server{
		Addr:    s.addr,
		Handler: midldewares(router),
	}

	log.Print("Server is running on ", s.addr)

	return server.ListenAndServe()

}
