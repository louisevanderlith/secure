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

func ProfilesSearch(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)

	db := core.Context()
	result, err := db.Profiles.Find(page, size, husk.Everything())

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ProfileView(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

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

	err = mix.Write(w, mix.JSON(result.Data()))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ProfileCreate(w http.ResponseWriter, r *http.Request) {
	body := prime.Profile{}
	err := drx.JSONBody(r, &body)

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

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ProfileUpdate(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := prime.Profile{}
	err = drx.JSONBody(r, &body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db := core.Context()
	db.UpdateProfile(key, body)

	if err != nil {
		log.Println("Update Profile Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
