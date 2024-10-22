package auth

import (
	"matwa/blogger/server/handlers/core"
	"net/http"

	"github.com/markbates/goth/gothic"
)

var Logout core.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}
