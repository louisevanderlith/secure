package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/secure/core"
)

type RegisterController struct {
	control.APIController
}

func NewRegisterCtrl(ctrlMap *control.ControllerMap) *RegisterController {
	result := &RegisterController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title Register
// @Description Registers a new user
// @Param	body		body 	core.AuthRequest		true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *RegisterController) Post() {
	var regis core.Registration
	json.Unmarshal(req.Ctx.Input.RequestBody, &regis)

	result, err := core.Register(regis)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
	}

	req.Serve(http.StatusOK, nil, result)
}
