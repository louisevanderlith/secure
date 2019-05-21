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

	lognCtrl := controllers.NewLoginCtrl(ctrlmap, privateKeyPath)

	beego.Router("/v1/login", lognCtrl, "post:Post")
	//beego.Router("/v1/login/:sessionID", lognCtrl, "delete:Logout")
	//beego.Router("/v1/login/avo/:sessionID", lognCtrl, "get:GetAvo")

	beego.Router("/v1/register", controllers.NewRegisterCtrl(ctrlmap), "post:Post")
	beego.Router("/v1/user/all/:pagesize", controllers.NewUserCtrl(ctrlmap), "get:Get")
}

func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(core.ActionMap)

	ctrlmap.Add("/v1/login", emptyMap)
	ctrlmap.Add("/v1/register", emptyMap)

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
