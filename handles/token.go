package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

func TokenPOST(w http.ResponseWriter, r *http.Request) {
	clnt, pass, ok := r.BasicAuth()

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

	require, err := req.GetRequirements()

	if err != nil {
		log.Println("Get Requirements Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	clms, err := Security.RequestToken(clnt, pass, req.Token, require)

	if err != nil {
		log.Println("Request Token Error", err)
		http.Error(w, "", http.StatusUnprocessableEntity)
		return
	}

	tkn, err := Security.Sign(clms, 5)

	if err != nil {
		log.Println("Sign Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(tkn))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
