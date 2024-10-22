package links

import (
	"matwa/blogger/server/handlers/core"
	"matwa/blogger/templates"
	"net/http"
)

var LoginLink core.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	LoginPage := templates.LoginPage()
	LoginPage.Render(r.Context(), w)
}
