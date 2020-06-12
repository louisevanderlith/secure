package handles

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func WhitelistGET(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	scp, pass, ok := r.BasicAuth()

	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	rsrc, err := Author.GetStore().GetResource(scp)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	if !rsrc.VerifySecret(pass) {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	res := Author.GetStore().GetWhitelist()

	err = ctx.Serve(http.StatusOK, mix.JSON(res))

	if err != nil {
		log.Println(err)
	}
}
