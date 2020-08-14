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

	clms, err := Security.RequestToken(clnt, pass, req.Token, req.Claims)

	if err != nil {
		log.Println("Request Token Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	tkn, err := Security.Sign(clms, 5)

	if err != nil {
		log.Println("Sign Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tkn))
}
