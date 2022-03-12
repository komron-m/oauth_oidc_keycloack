package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

type httpHandler struct {
	repo *dummyRepo
}

func (h *httpHandler) createHero(w http.ResponseWriter, r *http.Request) {
	heroReq := new(hero)
	if err := h.unmarshal(r, heroReq); err != nil {
		h.handlerError(w, r, err)
	}

	heroReq.ID = uuid.NewString()
	if err := h.repo.create(heroReq); err != nil {
		h.handlerError(w, r, err)
	}
	h.marshal(w, heroReq)
}

func (h *httpHandler) deleteHero(w http.ResponseWriter, r *http.Request) {
	type deleteHeroReq struct {
		Id string `json:"id"`
	}
	req := new(deleteHeroReq)
	if err := h.unmarshal(r, req); err != nil {
		h.handlerError(w, r, err)
	}

	if err := h.repo.delete(req.Id); err != nil {
		h.handlerError(w, r, err)
	}
}

func (h *httpHandler) getAllHeroes(w http.ResponseWriter, _ *http.Request) {
	heroes := h.repo.getAllHeroes()
	h.marshal(w, heroes)
}

func (h *httpHandler) unmarshal(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (h *httpHandler) marshal(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func (h *httpHandler) handlerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("URL: %s\n", r.URL.String())
	log.Printf("METHOD: %s\n", r.Method)
	log.Printf("ERROR: %+v\n", err)
	log.Println("--------------------------------------------------------")

	http.Error(w, err.Error(), http.StatusInternalServerError)
}
