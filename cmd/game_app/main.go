package main

import (
	"log"
	"net/http"
)

const (
	clientID = ""
	issuer   = ""
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

	mux.HandleFunc("/create", handler.createHero)
	mux.HandleFunc("/delete", handler.deleteHero)
	mux.HandleFunc("/get_all", handler.getAllHeroes)

	//oidcMid, err := internal.NewOpenIDCMid(issuer, clientID)
	//if err != nil {
	//	log.Fatal(err)
	//}

	err := http.ListenAndServe(":4000", cors(mux))
	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}
}
