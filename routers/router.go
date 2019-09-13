package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/secure/controllers"
)

func Setup(e resins.Epoxi, privateKey string) {
	forgotCtrl := &controllers.Forgot{}
	regCtrl := &controllers.Register{}

	loginCtrl := &controllers.Login{
		PrivateKey: privateKey,
	}

	e.JoinBundle("/", roletype.Unknown, mix.JSON, forgotCtrl, regCtrl, loginCtrl)

	usrCtrl := &controllers.User{}
	e.JoinBundle("/", roletype.Admin, mix.JSON, usrCtrl)
}
