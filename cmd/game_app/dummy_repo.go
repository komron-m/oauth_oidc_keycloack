package main

import (
	"sync"
)

type hero struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type dummyRepo struct {
	heroes []*hero
	sync.RWMutex
}

func (r *dummyRepo) create(h *hero) error {
	r.Lock()
	defer r.Unlock()

	r.heroes = append(r.heroes, h)
	return nil
}

func (r *dummyRepo) delete(id string) error {
	r.Lock()
	defer r.Unlock()

	for i := range r.heroes {
		if r.heroes[i].ID == id {
			r.heroes = append(r.heroes[:i], r.heroes[i+1:]...)
			break
		}
	}
	return nil
}

func (r *dummyRepo) getAllHeroes() []*hero {
	return r.heroes
}
