package handles

import (
	"encoding/json"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

func ConsentQuery(w http.ResponseWriter, r *http.Request) {
	obj := prime.QueryRequest{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := Security.QueryClient(obj.Token)

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
