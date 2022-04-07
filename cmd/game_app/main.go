package main

import (
	"context"
	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"log"
	"net/http"
	"strings"
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

type scopes struct {
	Scopes string `json:"scope"`
}

func (s *scopes) Validate(ctx context.Context) error {
	return nil
}

func (s *scopes) HasScope(scope string) bool {
	givenScopes := strings.Split(s.Scopes, " ")
	for _, gs := range givenScopes {
		if gs == scope {
			return true
		}
	}
	return false
}

func accessControlMid(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

		claims, ok := token.CustomClaims.(*scopes)
		if !ok || !claims.HasScope("openid") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message":"Insufficient scope."}`))
			return
		}

		next.ServeHTTP(w, r)
	}
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

	// uncomment and apply this + accessControlMid -- FOR OAuth2 and scopes check
	//oauthMid, err := internal.NewAccessTokenMid(issuer, audience, validator.ES256, new(scopes))

	err := http.ListenAndServe(":4000", cors(mux))
	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}
}
