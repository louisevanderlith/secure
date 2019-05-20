package core

import (
	"encoding/json"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/husk"
)

//Cookies is our Cookie object.
type Cookies struct {
	UserKey    husk.Key
	Username   string
	UserRoles  ActionMap
	IP         string
	Location   string
	Issuer     string    `json:"iss"`
	Audience   string    `json:"aud"`
	Expiration time.Time `json:"exp"`
	IssuedAt   time.Time `json:"iat"`
}

//NewCookies returns some new Cookies.
func NewCookies(userkey husk.Key, username, ip, location string, roles ActionMap) *Cookies {
	return &Cookies{
		UserKey:    userkey,
		Username:   username,
		IP:         ip,
		Location:   location,
		UserRoles:  roles,
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(time.Hour * 6),
		Issuer:     "https://secure.localhost/oauth/",
		Audience:   "https://localhost",
	}
}

//GetClaims return the JWT Claims from the Cookies Object
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
