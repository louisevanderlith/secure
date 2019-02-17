package main

import (
	"log"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	_ "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/routers"

	"github.com/astaxie/beego"
)

func main() {
	mode := beego.BConfig.RunMode

	if mode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
