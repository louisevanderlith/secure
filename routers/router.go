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
	"github.com/louisevanderlith/secure/logic"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/mango/enums"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	lognCtrl := controllers.NewLoginCtrl(ctrlmap)

	beego.Router("/v1/login", lognCtrl, "get:Get;post:Post")
	beego.Router("/v1/login/logout/:sessionID", lognCtrl, "get:Logout")
	beego.Router("/v1/login/avo/:sessionID", lognCtrl, "get:GetAvo")

	beego.Router("/v1/register", controllers.NewRegisterCtrl(ctrlmap), "get:Get;post:Post")
	beego.Router("/v1/user/all/:pagesize", controllers.NewUserCtrl(ctrlmap), "get:Get")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := logic.NewMasterMap(s)

	emptyMap := make(control.ActionMap)

	ctrlmap.Add("/login", emptyMap)
	ctrlmap.Add("/register", emptyMap)

	userMap := make(control.ActionMap)
	userMap["GET"] = enums.Admin

	ctrlmap.Add("/user", userMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterMaster)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap.ControllerMap
}
