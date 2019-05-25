package controllers

import (
	"net/http"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/secure/core"
)

type UserController struct {
	control.APIController
}

func NewUserCtrl(ctrlMap *control.ControllerMap) *UserController {
	result := &UserController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetUsers
// @Description Gets all Users
// @Success 200 {[]logic.UserObject]} []logic.UserObject]
// @router /all/:pagesize [get]
func (req *UserController) Get() {
	page, size := req.GetPageData()
	result := core.GetUsers(page, size)

	req.Serve(http.StatusOK, nil, result)
}

// @Title GetUser
// @Description Gets registered user
// @Param	key			path	string 	true		"User Key"
// @Success 200 {core.User} core.User
// @router /:key [get]
func (req *UserController) GetOne() {
	siteParam := req.Ctx.Input.Param(":key")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result, err := core.GetUser(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}
