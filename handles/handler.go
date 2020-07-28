package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/secure/core"
	"github.com/rs/cors"
	"net/http"
)

var Author kong.Author

func SetupRoutes(scrt string) http.Handler {
	authr, err := kong.CreateAuthority(core.Context(), )

	if err != nil {
		panic(err)
	}

	Author = authr

	r := mux.NewRouter()

	r.HandleFunc("/token", TokenPOST).Methods(http.MethodPost)

	r.HandleFunc("/login", kong.InternalMiddleware(authr, "kong.login.apply", scrt, LoginPOST)).Methods(http.MethodPost)
	r.HandleFunc("/consent", kong.InternalMiddleware(authr, "kong.consent.apply", scrt, ConsentPOST)).Methods(http.MethodPost)
	r.HandleFunc("/query", kong.InternalMiddleware(authr, "kong.client.query", scrt, ConsentQuery)).Methods(http.MethodPost)
	r.HandleFunc("/register", kong.InternalMiddleware(authr, "kong.user.register", scrt, RegisterPOST)).Methods(http.MethodPost)

	r.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", kong.InternalMiddleware(authr, "secure.resource.view", scrt, ResourcesView)).Methods(http.MethodGet)

	srchR := kong.InternalMiddleware(authr, "secure.resource.search", scrt, ResourcesSearch)
	r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}", srchR).Methods(http.MethodGet)
	//r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchR).Methods(http.MethodGet)

	r.HandleFunc("/resources", kong.InternalMiddleware(authr, "secure.resource.create", scrt, ResourcesCreate)).Methods(http.MethodPost)
	r.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", kong.InternalMiddleware(authr, "secure.resource.update", scrt, ResourcesUpdate)).Methods(http.MethodPut)

	r.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", kong.InternalMiddleware(authr, "secure.profile.view", scrt, ProfileView)).Methods(http.MethodGet)

	srchP := kong.InternalMiddleware(authr, "secure.profile.search", scrt, ProfilesSearch)
	r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}", srchP).Methods(http.MethodGet)
	//r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	r.HandleFunc("/profiles", kong.InternalMiddleware(authr, "secure.profile.create", scrt, ProfileCreate)).Methods(http.MethodPost)
	r.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", kong.InternalMiddleware(authr, "secure.profile.update", scrt, ProfileUpdate)).Methods(http.MethodPut)

	r.HandleFunc("/inspect", InspectPOST).Methods(http.MethodPost)
	r.HandleFunc("/info", InfoPOST).Methods(http.MethodPost)
	r.HandleFunc("/whitelist", WhitelistGET).Methods(http.MethodGet)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: core.Context().GetWhitelist(), //you service is available and allowed for this base url
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
