package main

import (
	"os"
	"path"
	"strconv"

	"github.com/louisevanderlith/secure/controllers"

	"github.com/louisevanderlith/droxolite/security/roletype"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/security/models"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	//pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	host := os.Getenv("HOST")
	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
	//appName := os.Getenv("APPNAME")
	//pubPath := path.Join(keyPath, pubName)
	privPath := path.Join(keyPath, privName)

	// Register with router
	/*	srv := bodies.NewService(appName, "", pubPath, host, httpport, "API")

		routr, err := do.GetServiceURL("", "Router.API", false)

		if err != nil {
			panic(err)
		}

		err = srv.Register(routr)

		if err != nil {
			panic(err)
		}*/

	thm := element.GetNoTheme(host, "mango.auth", "none")
	err := thm.LoadTemplate("./views", "master.html")

	if err != nil {
		panic(err)
	}

	poxy := resins.NewColourEpoxy(models.ClientCred{}, nil, element.GetNoTheme(host, "mango.auth", "none"), "secure.localhost", roletype.Unknown, controllers.Index)
	routers.Setup(poxy, privPath)
	poxy.EnableCORS(host)

	core.CreateContext()
	defer core.Shutdown()

	err = droxolite.Boot(poxy, httpport)

	if err != nil {
		panic(err)
	}
}
