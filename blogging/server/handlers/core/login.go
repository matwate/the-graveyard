package core

import (
	"matwa/blogger/data"
	"net/http"
)

var Login HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	password := r.Header.Get("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username and password are required"))
		return
	}

	for _, user := range data.UsersStore {
		if user.Username == username && user.Password == password {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(user.Token))
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Invalid credentials"))
}
