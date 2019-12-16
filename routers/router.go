package routers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/secure/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-session/session"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/secure/core"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func Setup(e resins.Epoxi, privateKey string) {
	srv := SetupOAuthServer()
	authCtrl := controllers.NewOAuth2(srv)
	e.JoinPath(e.Router().(*mux.Router), "/auth", "Auth Endpoint", http.MethodPost, 0, mix.JSON, authCtrl.Auth)
	e.JoinPath(e.Router().(*mux.Router), "/authorize", "Authorise Endpoint", http.MethodPost, 0, mix.JSON, authCtrl.Authorize)
	e.JoinPath(e.Router().(*mux.Router), "/token", "Token Endpoint", http.MethodGet, 0, mix.JSON, authCtrl.Token)

	lognCtrl := &controllers.Login{}
	e.JoinBundle("/", 0, mix.Page, lognCtrl)

	/*forgotCtrl := &controllers.Forgot{}
	regCtrl := &controllers.Register{}

	loginCtrl := &controllers.Login{
		PrivateKey: privateKey,
	}

	e.JoinBundle("/", 0, mix.JSON, forgotCtrl, regCtrl, loginCtrl)

	usrCtrl := &controllers.User{}
	e.JoinBundle("/", 1, mix.JSON, usrCtrl)*/
}

func SetupOAuthServer() *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	/*clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://localhost:9094",
	})*/

	manager.MapClientStorage(core.NewClientStore())

	srv := server.NewServer(server.NewConfig(), manager)

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return srv
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
