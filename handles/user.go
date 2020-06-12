package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/secure/core"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	result := core.GetUsers(1, 10)

	err := ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}

// @Title GetUsers
// @Description Gets all Users
// @Success 200 {[]logic.UserObject]} []logic.UserObject]
// @router /all/:pagesize [get]
func SearchUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()
	result := core.GetUsers(page, size)

	err := ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}

// @Title GetUser
// @Description Gets registered user
// @Param	key			path	string 	true		"User Key"
// @Success 200 {core.User} core.User
// @router /:key [get]
func ViewUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	siteParam := ctx.FindParam("key")

	_, err := husk.ParseKey(siteParam)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result := core.Context().GetUser(siteParam)

	err = ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}

// @router /:key [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	siteParam := ctx.FindParam("key")

	_, err := husk.ParseKey(siteParam)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var roles []core.Role
	err = ctx.Body(&roles)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	//err = core.UpdateRoles(key, roles)

	//if err != nil {
	//	return http.StatusInternalServerError, err
	//}

	err = ctx.Serve(http.StatusOK, mix.JSON("Updated User"))

	if err != nil {
		log.Println(err)
	}
}
