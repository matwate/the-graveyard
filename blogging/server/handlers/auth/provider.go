package auth

import (
	"matwa/blogger/server/handlers/core"
	"net/http"
)

var Provider core.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
	provider := req.PathValue("provider")

	if provider == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Provider not found"))
	}

	q := req.URL.Query()
	q.Add("provider", provider)
	req.URL.RawQuery = q.Encode()

	if 

}
