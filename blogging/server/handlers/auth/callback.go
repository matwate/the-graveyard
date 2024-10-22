package auth

import (
	"context"
	"fmt"
	"matwa/blogger/server/handlers/core"
	"net/http"

	"github.com/markbates/goth/gothic"
)

var HandleCallback core.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

	prov := r.PathValue("provider")

	r = r.WithContext(context.WithValue(r.Context(), "provider", prov))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("User: %#v", user)))

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
