package logic

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/secure/core"
)

// AttemptLogin returns SessionID, if error is not nil
func AttemptLogin(ctx *context.Context) (string, error) {
	authReq := core.Authentication{}
	err := json.Unmarshal(ctx.Input.RequestBody, &authReq)

	if err != nil {
		return "", err
	}

	cooki, err := core.Login(authReq)

	if err != nil {
		return "", err
	}

	sessionID := CreateAvo(ctx, cooki)

	return sessionID, nil
}
