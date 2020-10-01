package handles

import (
	"encoding/json"
	"github.com/louisevanderlith/droxolite/drx"
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

	bits, err := json.Marshal(res)

	if err != nil {
		log.Println("Marshal Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(bits)
}
