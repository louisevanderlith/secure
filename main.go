package main

import (
	"log"
	"os"
	"path"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/routers"

	"github.com/astaxie/beego"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)
	privPath := path.Join(keyPath, privName)
	
	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(appName, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv, privPath, host)

		beego.Run()
	}
}
