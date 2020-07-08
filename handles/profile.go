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

func ProfilesSearch(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	page, size := ctx.GetPageData()

	db := core.Context()
	result, err := db.Profiles.Find(page, size, husk.Everything())

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result))
}

func ProfileView(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	result, err := db.Profiles.FindByKey(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(result.Data()))
}

func ProfileCreate(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	body := prime.Profile{}
	err := ctx.Body(&body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	rec, err := db.Profiles.Create(body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))
}

func ProfileUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := prime.Profile{}
	err = ctx.Body(&body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	db.UpdateProfile(key, body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(nil))
}
