package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
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

	alg := jwt.GetSigningMethod("RS256")
	if alg == nil {
		return "", fmt.Errorf("Couldn't find signing method: %v", "RS256")
	}

	token := jwt.NewWithClaims(alg, cooki.GetClaims())

	var rdr io.Reader
	if f, err := os.Open("/db/sign_rsa"); err == nil {
		rdr = f
		defer f.Close()
	} else {
		return "", err
	}

	bits, err := ioutil.ReadAll(rdr)

	if err != nil {
		return "", err
	}

	pkey, err := jwt.ParseRSAPrivateKeyFromPEM(bits)

	if err != nil {
		return "", err
	}

	result, err := token.SignedString(pkey)

	return result, err
}
