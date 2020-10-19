package main

import (
	"flag"
	"github.com/louisevanderlith/secure/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/secure/core"
)

func main() {
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8086",
		Handler:      handles.SetupRoutes(*srcSecrt),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
