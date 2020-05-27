package handles

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/secure/core"
)

type User struct {
}

func (req *User) Get(ctx context.Requester) (int, interface{}) {
	result := core.GetUsers(1, 10)

	return http.StatusOK, result
}

// @Title GetUsers
// @Description Gets all Users
// @Success 200 {[]logic.UserObject]} []logic.UserObject]
// @router /all/:pagesize [get]
func (req *User) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()
	result := core.GetUsers(page, size)

	return http.StatusOK, result
}

// @Title GetUser
// @Description Gets registered user
// @Param	key			path	string 	true		"User Key"
// @Success 200 {core.User} core.User
// @router /:key [get]
func (req *User) View(ctx context.Requester) (int, interface{}) {
	siteParam := ctx.FindParam("key")

	_, err := husk.ParseKey(siteParam)

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := core.Context().GetUser(siteParam)

	return http.StatusOK, result
}

// @router /:key [put]
func (req *User) Update(ctx context.Requester) (int, interface{}) {
	siteParam := ctx.FindParam("key")

	_, err := husk.ParseKey(siteParam)

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	var roles []core.Role
	err = ctx.Body(&roles)

	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}

	//err = core.UpdateRoles(key, roles)

	//if err != nil {
	//	return http.StatusInternalServerError, err
	//}

	return http.StatusOK, "Updated Roles"
}
