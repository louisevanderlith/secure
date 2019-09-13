package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/secure/core"
)

type RegisterController struct {
	xontrols.APICtrl
}

// @Title Register
// @Description Registers a new user
// @Param	body		body 	core.AuthRequest		true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *RegisterController) Post() {
	var regis core.Registration
	err := req.Body(&regis)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
	}

	result, err := core.Register(regis)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}
