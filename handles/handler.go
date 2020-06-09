package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/secure/core"
	"github.com/rs/cors"
	"net/http"
)

var Author kong.Author

func SetupRoutes() http.Handler {
	authr, err := kong.CreateAuthority(core.Context())

	if err != nil {
		panic(err)
	}

	Author = authr

	r := mux.NewRouter()

	r.HandleFunc("/token", TokenPOST).Methods(http.MethodPost)

	r.HandleFunc("/login", kong.InternalMiddleware(authr, "kong.login.apply", "secret", LoginPOST)).Methods(http.MethodPost)
	r.HandleFunc("/consent", kong.InternalMiddleware(authr, "kong.consent.apply", "secret", ConsentPOST)).Methods(http.MethodPost)
	r.HandleFunc("/query", kong.InternalMiddleware(authr, "kong.client.query", "secret", ConsentQuery)).Methods(http.MethodPost)

	r.HandleFunc("/inspect", InspectPOST).Methods(http.MethodPost)
	r.HandleFunc("/info", InfoPOST).Methods(http.MethodPost)

	white := core.Context().GetWhitelist()

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: white, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
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
