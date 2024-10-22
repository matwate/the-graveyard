package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

const (
	key    = "testKey"
	MaxAge = 3600
	IsProd = false
)

func NewAuth() {
	githubCLientId := os.Getenv("GITHUB_KEY")
	githubClientSecret := os.Getenv("GITHUB_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		github.New(githubCLientId, githubClientSecret, "http://localhost:8800/auth/github/callback"),
	)
}
