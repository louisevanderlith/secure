package handles

import (
	"encoding/json"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

func InfoPOST(w http.ResponseWriter, r *http.Request) {
	_, pass, ok := r.BasicAuth()

	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	req := prime.QueryRequest{}
	err := drx.JSONBody(r, &req)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	claims, err := Security.ClientInsight(req.Token, pass)

	if err != nil {
		log.Println("Info Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	bits, err := json.Marshal(claims)

	if err != nil {
		log.Println("Marshal Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(bits)
}
