package handles

import (
	"encoding/json"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

func TokenPOST(w http.ResponseWriter, r *http.Request) {
	clnt, pass, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(nil)
		return
	}

	dec := json.NewDecoder(r.Body)
	req := prime.TokenReq{}
	err := dec.Decode(&req)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	clms, err := Author.RequestToken(clnt, pass, req.UserToken, req.Scopes...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tkn, err := Author.Sign(clms)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tkn))
}
