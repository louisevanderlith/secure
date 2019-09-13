package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	host := os.Getenv("HOST")
	profile := os.Getenv("PROFILE")
	pubPath := path.Join(keyPath, pubName)
	privPath := path.Join(keyPath, privName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Register with router
	srv := bodies.NewService(conf.Appname, pubPath, conf.HTTPPort, servicetype.API)

	routr, err := do.GetServiceURL("", "Router.API", false)

	if err != nil {
		panic(err)
	}

	err = srv.Register(routr)

	if err != nil {
		panic(err)
	}

	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, profile))
	routers.Setup(poxy, privPath)
	poxy.EnableCORS(host)

	core.CreateContext()
	defer core.Shutdown()

	err = droxolite.Boot(poxy)

	if err != nil {
		panic(err)
	}
}
