package handles

import (
	"encoding/json"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

func InspectPOST(w http.ResponseWriter, r *http.Request) {
	scp, pass, ok := r.BasicAuth()

	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	dec := json.NewDecoder(r.Body)
	req := prime.InspectReq{}
	err := dec.Decode(&req)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	claims, err := Author.Inspect(req.AccessCode, scp, pass)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	bits, err := json.Marshal(claims)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(bits)
}
