package logic

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/secure/core"
)

// AttemptLogin returns SessionID, if error is not nil
func AttemptLogin(ctx context.Requester, fullKeyPath string) (string, error) {
	authReq := core.Authentication{}
	err := ctx.Body(&authReq)

	if err != nil {
		return "", err
	}

	/*cooki, err := core.Login(authReq)

	if err != nil {
		return "", err
	}*/

	alg := jwt.GetSigningMethod("RS256")
	if alg == nil {
		return "", fmt.Errorf("Couldn't find signing method: %v", "RS256")
	}

	token := jwt.NewWithClaims(alg, nil)

	bits, err := readPrivateKey(fullKeyPath)

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

func readPrivateKey(fullKeyPath string) ([]byte, error) {
	var rdr io.Reader
	log.Println(fullKeyPath)
	if f, err := os.Open(fullKeyPath); err == nil {
		rdr = f
		defer f.Close()
	} else {
		return nil, err
	}

	return ioutil.ReadAll(rdr)
}
