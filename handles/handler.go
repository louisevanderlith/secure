package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/kong/middle"
	"github.com/louisevanderlith/secure/core"
	"github.com/rs/cors"
	"net/http"
)

var Security middle.Security

func SetupRoutes(scrt string) http.Handler {
	sec, err := kong.CreateSecurity(core.Context())

	if err != nil {
		panic(err)
	}

	Security = sec

	r := mux.NewRouter()

	r.HandleFunc("/token", TokenPOST).Methods(http.MethodPost)
	r.HandleFunc("/query/{client:[a-z]+\\.[a-z]+}", middle.InternalMiddleware(sec, "secure.client.query", scrt, ConsentQuery)).Methods(http.MethodGet)

	r.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", middle.InternalMiddleware(sec, "secure.resource.view", scrt, ResourcesView)).Methods(http.MethodGet)

	srchR := middle.InternalMiddleware(sec, "secure.resource.search", scrt, ResourcesSearch)
	r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}", srchR).Methods(http.MethodGet)
	//r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchR).Methods(http.MethodGet)

	r.HandleFunc("/resources", middle.InternalMiddleware(sec, "secure.resource.create", scrt, ResourcesCreate)).Methods(http.MethodPost)
	r.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", middle.InternalMiddleware(sec, "secure.resource.update", scrt, ResourcesUpdate)).Methods(http.MethodPut)

	r.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", middle.InternalMiddleware(sec, "secure.profile.view", scrt, ProfileView)).Methods(http.MethodGet)

	srchP := middle.InternalMiddleware(sec, "secure.profile.search", scrt, ProfilesSearch)
	r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}", srchP).Methods(http.MethodGet)
	//r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	r.HandleFunc("/profiles", middle.InternalMiddleware(sec, "secure.profile.create", scrt, ProfileCreate)).Methods(http.MethodPost)
	r.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", middle.InternalMiddleware(sec, "secure.profile.update", scrt, ProfileUpdate)).Methods(http.MethodPut)

	r.HandleFunc("/inspect", InspectPOST).Methods(http.MethodPost)
	r.HandleFunc("/info", InfoPOST).Methods(http.MethodPost)
	r.HandleFunc("/whitelist", WhitelistGET).Methods(http.MethodGet)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: core.Context().GetWhitelist("secure"), //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
