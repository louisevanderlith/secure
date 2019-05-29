// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/secure/controllers"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
)

func Setup(s *mango.Service, privateKeyPath, host string) {
	ctrlmap := EnableFilter(s, host)

	beego.Router("/v1/login", controllers.NewLoginCtrl(ctrlmap, privateKeyPath), "post:Post")

	beego.Router("/v1/register", controllers.NewRegisterCtrl(ctrlmap), "post:Post")

	usrCtrl := controllers.NewUserCtrl(ctrlmap)
	beego.Router("/v1/user/all/:pagesize", usrCtrl, "get:Get")
	beego.Router("/v1/user/:key", usrCtrl, "get:GetOne")

	forgetCtrl := controllers.NewForgotCtrl(ctrlmap)
	beego.Router("/v1/forgot/:forgotKey", forgetCtrl, "get:Get")
	beego.Router("/v1/forgot", forgetCtrl, "post:Post")
}

func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(core.ActionMap)

	ctrlmap.Add("/v1/login", emptyMap)
	ctrlmap.Add("/v1/register", emptyMap)
	ctrlmap.Add("/v1/forgot")

	userMap := make(core.ActionMap)
	userMap["GET"] = roletype.Admin

	ctrlmap.Add("/v1/user", userMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	return ctrlmap
}
