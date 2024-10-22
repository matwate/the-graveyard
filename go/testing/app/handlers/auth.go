package handlers

import (
	"testing/app/types"

	"github.com/anthdm/superkit/kit"
)

func HandleAuthentication(kit *kit.Kit) (kit.Auth, error) {
	return types.AuthUser{}, nil
}
