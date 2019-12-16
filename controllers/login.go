package controllers

import (
	"net/http"

	"github.com/go-session/session"
	"github.com/louisevanderlith/droxolite/context"
)

type Login struct {
}

func (x *Login) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusOK, nil
}

func (req *Login) Search(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *Login) Create(ctx context.Requester) (int, interface{}) {
	store, err := session.Start(nil, ctx.Responder(), ctx.Request())

	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	r := ctx.Request()

	if r.Form == nil {
		if err := r.ParseForm(); err != nil {
			return http.StatusInternalServerError, err.Error()
		}
	}
	store.Set("LoggedInUserID", r.Form.Get("username"))
	store.Save()

	//w.Header().Set("Location", "/auth")
	ctx.Redirect(http.StatusFound, "/auth")
	return http.StatusFound, nil
}
