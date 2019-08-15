package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/secure/core"
)

type Forgot struct {
	xontrols.APICtrl
}

func (req *Forgot) Get() {

}

// @Title Forgot Password
// @Description Will send the user an email with an OTP
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *Forgot) Post() {
	email := ""
	err := req.Body(&email)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
	}

	resp, err := core.RequestReset(email, req.Ctx().RequestURI())

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, resp)
}
