package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/secure/core"
)

// @Title Register
// @Description Registers a new user
// @Param	body		body 	core.AuthRequest		true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	var regis core.Registration
	err := ctx.Body(&regis)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.Register(regis)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}
