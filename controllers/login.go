package controllers

import (
	"errors"

	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/secure/logic"
)

type LoginController struct {
	control.APIController
}

func NewLoginCtrl(ctrlMap *control.ControllerMap) *LoginController {
	result := &LoginController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetAvo
// @Description Gets the currently logged in user's avo
// @Param	path	path	string	true	"sessionID"
// @Success 200 {map[string]string} map[string]string
// @router /avo/:sessionID [get]
func (req *LoginController) GetAvo() {
	sessionID := req.Ctx.Input.Param(":sessionID")
	hasAvo := logic.HasAvo(sessionID)

	if !hasAvo {
		req.Serve(nil, errors.New("no avo found"))
		return
	}

	result := logic.FindAvo(sessionID)

	req.Serve(result, nil)
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *LoginController) Post() {
	sessionID, err := logic.AttemptLogin(req.Ctx)

	req.Serve(sessionID, err)
}

// @Title Logout
// @Description Logs out current logged in user session
// @Param	path	path	string	true	"sessionID"
// @Success 200 {string} string
// @router /logout/:sessionID [get]
func (req *LoginController) Logout() {
	sessionID := req.Ctx.Input.Param(":sessionID")

	// TODO: Create Trace for Logout...
	logic.DestroyAvo(sessionID)

	req.Serve("Logout Success", nil)
}
