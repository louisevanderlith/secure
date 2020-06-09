package handles

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func WhitelistGET(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	//res := Author.GetStore().GetWhitelist()

	err := ctx.Serve(http.StatusOK, mix.JSON(nil))

	if err != nil {
		log.Println(err)
	}
}
