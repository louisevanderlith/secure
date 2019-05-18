// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/secure/controllers"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
)

func Setup(s *mango.Service, privateKeyPath string) {
	ctrlmap := EnableFilter(s)

	lognCtrl := controllers.NewLoginCtrl(ctrlmap, privateKeyPath)

	beego.Router("/v1/login", lognCtrl, "post:Post")
	//beego.Router("/v1/login/:sessionID", lognCtrl, "delete:Logout")
	//beego.Router("/v1/login/avo/:sessionID", lognCtrl, "get:GetAvo")

	beego.Router("/v1/register", controllers.NewRegisterCtrl(ctrlmap), "post:Post")
	beego.Router("/v1/user/all/:pagesize", controllers.NewUserCtrl(ctrlmap), "get:Get")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(core.ActionMap)

	ctrlmap.Add("/v1/login", emptyMap)
	ctrlmap.Add("/v1/register", emptyMap)

	userMap := make(core.ActionMap)
	userMap["GET"] = roletype.Admin

	ctrlmap.Add("/v1/user", userMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
