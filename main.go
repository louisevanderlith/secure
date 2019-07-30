package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	//host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)
	privPath := path.Join(keyPath, privName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Register with router
	srv := droxolite.NewService(conf.Appname, pubPath, conf.HTTPPort, servicetype.API)

	err = srv.Register()

	if err != nil {
		panic(err)
	}

	poxy := droxolite.NewEpoxy(srv)
	routers.Setup(poxy, privPath)

	core.CreateContext()
	defer core.Shutdown()

	err = poxy.Boot()

	if err != nil {
		panic(err)
	}
}
