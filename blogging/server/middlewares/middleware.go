package middlewares

import (
	"matwa/blogger/server/middlewares/core"
)

var ApplyMiddlewareFuncs = core.CreateStack(
	core.Logging,
)
