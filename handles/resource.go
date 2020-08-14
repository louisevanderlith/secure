package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
	"github.com/louisevanderlith/secure/core"
	"log"
	"net/http"
)

func ResourcesSearch(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)

	db := core.Context()
	result, err := db.Resources.Find(page, size, husk.Everything())

	if err != nil {
		log.Println("Find Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(result))
}

func ResourcesView(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	result, err := db.Resources.FindByKey(key)

	if err != nil {
		log.Println("Find Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(result.Data()))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ResourcesCreate(w http.ResponseWriter, r *http.Request) {
	body := prime.Resource{}
	err := drx.JSONBody(r, &body)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	rec, err := db.Resources.Create(body)

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(rec))
}

func ResourcesUpdate(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := prime.Resource{}
	err = drx.JSONBody(r, &body)

	if err != nil {
		log.Println("Bind Error", err)
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

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
