package controllers

import (
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
	req.Serve(result, nil)
}
