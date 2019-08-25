package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/secure/controllers"
)

func Setup(poxy resins.Epoxi, privateKey string) {
	//Forgot
	forgotCtrl := &controllers.Forgot{}
	forgtGroup := routing.NewRouteGroup("forgot", mix.JSON)
	forgtGroup.AddRoute("Forgot Request", "/{forgotKey:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, forgotCtrl.Get)
	forgtGroup.AddRoute("Create Request", "", "POST", roletype.Unknown, forgotCtrl.Post)
	poxy.AddGroup(forgtGroup)

	//Login
	loginCtrl := &controllers.Login{
		PrivateKey: privateKey,
	}

	lognGroup := routing.NewRouteGroup("login", mix.JSON)
	lognGroup.AddRoute("Create Login", "", "POST", roletype.Unknown, loginCtrl.Post)
	poxy.AddGroup(lognGroup)

	//Register
	regCtrl := &controllers.Register{}
	regGroup := routing.NewRouteGroup("register", mix.JSON)
	regGroup.AddRoute("Create Reqistration", "", "POST", roletype.Unknown, regCtrl.Post)
	poxy.AddGroup(regGroup)

	//User
	usrCtrl := &controllers.User{}
	usrGroup := routing.NewRouteGroup("user", mix.JSON)
	usrGroup.AddRoute("User by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, usrCtrl.GetOne)
	usrGroup.AddRoute("Update Roles", "/{key:[0-9]+\x60[0-9]+}", "PUT", roletype.Admin, usrCtrl.UpdateRoles)
	usrGroup.AddRoute("All Users", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, usrCtrl.Get)
	poxy.AddGroup(usrGroup)
	/*
		ctrlmap := EnableFilter(s, host)
		beego.Router("/v1/login", controllers.NewLoginCtrl(ctrlmap, privateKeyPath), "post:Post")

		beego.Router("/v1/register", controllers.NewRegisterCtrl(ctrlmap), "post:Post")

		usrCtrl := controllers.NewUserCtrl(ctrlmap)
		beego.Router("/v1/user/all/:pagesize", usrCtrl, "get:Get")
		beego.Router("/v1/user/:key", usrCtrl, "get:GetOne;put:UpdateRoles")

		forgetCtrl := controllers.NewForgotCtrl(ctrlmap)
		beego.Router("/v1/forgot/:forgotKey", forgetCtrl, "get:Get")
		beego.Router("/v1/forgot", forgetCtrl, "post:Post")*/
}

/*
func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(core.ActionMap)

	ctrlmap.Add("/v1/login", emptyMap)
	ctrlmap.Add("/v1/register", emptyMap)
	ctrlmap.Add("/v1/forgot", emptyMap)

	userMap := make(core.ActionMap)
	userMap["GET"] = roletype.Admin
	userMap["PUT"] = roletype.Admin

	ctrlmap.Add("/v1/user", userMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "PUT", "POST", "OPTIONS"},
	}))

	return ctrlmap
}
*/
