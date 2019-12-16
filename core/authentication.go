package core

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	App      Application
	Email    string
	Password string
}

// password hashing cost
const cost int = 11

// Login will attempt to authenticate a user
func Login(authReq Authentication) (interface{}, error) {

	if len(authReq.Password) < 6 {
		return nil, errors.New("password must be longer than 6 characters")
	}

	if !strings.Contains(authReq.Email, "@") {
		return nil, errors.New("email is invalid")
	}

	userRec, err := getUser(authReq.Email)

	if err != nil {
		return nil, err
	}

	user := userRec.Data().(User)

	if !user.Verified {
		return nil, errors.New("user not yet verified")
	}

	compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authReq.Password))
	passed := compare == nil
	user.AddTrace(getLoginTrace(authReq, passed))
	err = ctx.Users.Update(userRec)

	if err != nil {
		return nil, err
	}

	defer ctx.Users.Save()

	if !passed {
		return nil, errors.New("login failed")
	}

	return nil, nil
	//return bodies.NewCookies(userRec.GetKey(), user.Name, ip, location, user.Email, user.RoleMap()), nil
}
