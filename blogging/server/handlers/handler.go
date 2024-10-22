package handlers

import (
	"matwa/blogger/server/handlers/auth"
	"matwa/blogger/server/handlers/core"
	"matwa/blogger/server/handlers/links"
)

var Handlers = map[string]core.HandlerFunc{
	"login":      core.Login,
	"register":   core.Register,
	"newBlog":    core.NewBlog,
	"getBlog":    core.GetBlog,
	"updateBlog": core.UpdateBlog,
	"link-Login": links.LoginLink,
}

var AuthHandlers = map[string]core.HandlerFunc{
	"callback": auth.HandleCallback,
	"logout":   auth.Logout,
	"provider": auth.Provider,
}
