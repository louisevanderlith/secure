package controllers

import (
	"errors"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/secure/core"
)

type Forgot struct {
}

func (req *Forgot) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, errors.New("to do")
}

func (x *Forgot) Search(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

// @Title Forgot Password
// @Description Will send the user an email with an OTP
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *Forgot) Create(ctx context.Requester) (int, interface{}) {
	email := ""
	err := ctx.Body(&email)

	if err != nil {
		return http.StatusBadRequest, err
	}

	resp, err := core.RequestReset(email, ctx.RequestURI())

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, resp
}
