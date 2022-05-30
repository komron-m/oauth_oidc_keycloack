package main

import (
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/komron-m/oauth_oidc_keycloack/internal"
	"log"
	"net/http"
)

var (
	clientID = "finmonitoring"
	issuer   = "http://localhost:8080/realms/demo"
	audience []string
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	handler := &httpHandler{repo: new(dummyRepo)}

	mux.HandleFunc("/create", handler.create)
	mux.HandleFunc("/delete", handler.delete)
	mux.HandleFunc("/get_all", handler.getAll)

	// uncomment and apply this -- FOR OpenIdConnect
	//oidcMid, err := internal.NewOpenIDCMid(issuer, clientID)
	//if err != nil {
	//	log.Fatal(err)
	//}

	oauthMid, err := internal.NewAccessTokenMid(issuer, audience, validator.ES256, new(internal.Scope))
	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}

	err = http.ListenAndServe(":4000", oauthMid(cors(mux)))
	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}
}
