package core

import (
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
