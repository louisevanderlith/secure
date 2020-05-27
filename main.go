package main

import (
	"encoding/gob"
	"github.com/louisevanderlith/kong/tokens"
	"github.com/louisevanderlith/secure/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/secure/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	gob.Register(tokens.Claims{})

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8086",
		Handler:      handles.SetupRoutes(),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
