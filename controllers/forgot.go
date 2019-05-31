package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/secure/core"
)

type ForgotController struct {
	control.APIController
}

func NewForgotCtrl(ctrlMap *control.ControllerMap) *ForgotController {
	result := &ForgotController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (req *ForgotController) Get() {

}

// @Title Forgot Password
// @Description Will send the user an email with an OTP
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *ForgotController) Post() {
	email := ""
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &email)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
	}

	resp, err := core.RequestReset(email, req.Ctx.Request.URL.RequestURI())

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, resp)
}
