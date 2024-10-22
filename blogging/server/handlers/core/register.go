package core

import (
	"fmt"
	"matwa/blogger/data"
	"net/http"
)

var Register HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	password := r.Header.Get("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username and password are required"))
		return
	}

	data.UsersStore = append(data.UsersStore, data.User{
		Username: username,
		Password: password,
		Token:    fmt.Sprintf("%s:%s", username, password),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered"))
}
