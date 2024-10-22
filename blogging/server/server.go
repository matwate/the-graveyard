package server

import (
	"log"
	"matwa/blogger/server/auth"
	"matwa/blogger/server/handlers"
	m "matwa/blogger/server/middlewares"
	"matwa/blogger/templates"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() error {

	auth.NewAuth()

	router := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	HomePage := templates.HomePage("Galactik Blog")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomePage.Render(r.Context(), w)
	})

	router.HandleFunc("/login", handlers.Handlers["login"])
	router.HandleFunc("/register", handlers.Handlers["register"])
	router.HandleFunc("POST /NBlog", handlers.Handlers["newBlog"])
	router.HandleFunc("GET /Blog/{title}/{index}", handlers.Handlers["getBlog"])
	router.HandleFunc("PUT /Blog/{title}", handlers.Handlers["updateBlog"])

	//Htmx Handlers
	router.HandleFunc("GET /link-login", handlers.Handlers["link-Login"])
	funcChain := m.ApplyMiddlewareFuncs(router)

	//Auth Handlers
	router.HandleFunc("/auth/{provider}/callback", handlers.AuthHandlers["callback"])
	router.HandleFunc("logout/{provider}", handlers.AuthHandlers["logout"])
	router.HandleFunc("/auth/{provider}", handlers.AuthHandlers["provider"])

	server := &http.Server{
		Addr:    s.addr,
		Handler: funcChain,
	}

	log.Print("Server is running on ", s.addr)

	return server.ListenAndServe()

}
