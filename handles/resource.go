package handles

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
	"github.com/louisevanderlith/secure/core"
	"log"
	"net/http"
)

func ResourcesSearch(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	page, size := ctx.GetPageData()

	db := core.Context()
	result, err := db.Resources.Find(page, size, husk.Everything())

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))
}

func ResourcesView(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	result, err := db.Resources.FindByKey(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result.Data()))
}

func ResourcesCreate(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	body := prime.Resource{}
	err := ctx.Body(&body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	rec, err := db.Resources.Create(body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))
}

func ResourcesUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := prime.Resource{}
	err = ctx.Body(&body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	db.UpdateResource(key, body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(nil))
}
