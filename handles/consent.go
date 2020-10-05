package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func ConsentQuery(w http.ResponseWriter, r *http.Request) {
	client := drx.FindParam(r, "client")
	res, err := Security.ClientResourceQuery(client)

	if err != nil {
		log.Println("Query Client Error", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = mix.Write(w, mix.JSON(res))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
