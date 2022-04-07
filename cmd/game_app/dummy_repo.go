package main

import (
	"sync"
)

type user struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type dummyRepo struct {
	users []*user
	sync.RWMutex
}

func (r *dummyRepo) create(h *user) error {
	r.Lock()
	defer r.Unlock()

	r.users = append(r.users, h)
	return nil
}

func (r *dummyRepo) delete(id string) error {
	r.Lock()
	defer r.Unlock()

	for i := range r.users {
		if r.users[i].ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}
	return nil
}

func (r *dummyRepo) getAll() []*user {
	return r.users
}
