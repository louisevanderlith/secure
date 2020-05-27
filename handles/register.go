package handles

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/secure/core"
)

type Register struct {
}

func (x *Register) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

// @Title Register
// @Description Registers a new user
// @Param	body		body 	core.AuthRequest		true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *Register) Create(ctx context.Requester) (int, interface{}) {
	var regis core.Registration
	err := ctx.Body(&regis)

	if err != nil {
		return http.StatusBadRequest, err
	}

	result, err := core.Register(regis)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, result
}
