package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/secure/core"
)

type User struct {
	xontrols.APICtrl
}

// @Title GetUsers
// @Description Gets all Users
// @Success 200 {[]logic.UserObject]} []logic.UserObject]
// @router /all/:pagesize [get]
func (req *User) Get() {
	page, size := req.GetPageData()
	result := core.GetUsers(page, size)

	req.Serve(http.StatusOK, nil, result)
}

// @Title GetUser
// @Description Gets registered user
// @Param	key			path	string 	true		"User Key"
// @Success 200 {core.User} core.User
// @router /:key [get]
func (req *User) GetOne() {
	siteParam := req.FindParam("key")

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

// @router /:key [put]
func (req *User) UpdateRoles() {
	siteParam := req.FindParam("key")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	var roles []core.Role
	err = req.Body(&roles)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = core.UpdateRoles(key, roles)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, "Updated Roles")
}
