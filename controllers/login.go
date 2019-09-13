package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/secure/logic"
)

type LoginController struct {
	xontrols.APICtrl
	PrivateKey string
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *LoginController) Post() {
	sessionID, err := logic.AttemptLogin(req.Ctx(), req.PrivateKey)

	if err != nil {
		req.Serve(http.StatusForbidden, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, sessionID)
}
