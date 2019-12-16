package controllers

import (
	"net/http"
	"net/url"

	"gopkg.in/oauth2.v3/server"

	"github.com/go-session/session"
	"github.com/louisevanderlith/droxolite/context"
)

type OAuth2 struct {
	srv *server.Server
}

func NewOAuth2(srv *server.Server) *OAuth2 {
	return &OAuth2{srv}
}

func (o *OAuth2) Auth(ctx context.Requester) (int, interface{}) {
	store, err := session.Start(nil, ctx.Responder(), ctx.Request())

	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		//w.Header().Set("Location", "/login")
		//w.WriteHeader(http.StatusFound)

		ctx.Redirect(http.StatusFound, "/login")
		return http.StatusFound, nil
	}

	return http.StatusOK, nil
}

func (o *OAuth2) Authorize(ctx context.Requester) (int, interface{}) {
	w := ctx.Responder()
	r := ctx.Request()
	store, err := session.Start(nil, w, r)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}

	r.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = o.srv.HandleAuthorizeRequest(w, r)

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusOK, nil
}

func (o *OAuth2) Token(ctx context.Requester) (int, interface{}) {
	w := ctx.Responder()
	r := ctx.Request()
	err := o.srv.HandleTokenRequest(w, r)

	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, nil
}
