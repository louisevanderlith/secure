package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/secure/logic"
)

type Login struct {
	PrivateKey string
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *Login) Post(ctx context.Contexer) (int, interface{}) {
	sessionID, err := logic.AttemptLogin(ctx, req.PrivateKey)

	if err != nil {
		return http.StatusForbidden, err
	}

	return http.StatusOK, sessionID
}
