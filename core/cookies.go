package core

import (
	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/husk"
)

//Cookies is our Cookie object.
type Cookies struct {
	UserKey   husk.Key
	Username  string
	UserRoles ActionMap
	IP        string
	Location  string
}

//NewCookies returns some new Cookies.
func NewCookies(userkey husk.Key, username, ip, location string, roles ActionMap) *Cookies {
	return &Cookies{
		UserKey:   userkey,
		Username:  username,
		IP:        ip,
		Location:  location,
		UserRoles: roles,
	}
}

func (c Cookies) GetClaims() jwt.MapClaims {
	result := make(jwt.MapClaims)

	data, err := json.Marshal(c)

	if err != nil {
		return nil
	}

	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil
	}

	return result
}
