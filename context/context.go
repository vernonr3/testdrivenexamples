package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	//Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch(ctx context.Context) (string, error) {
	return s.response, nil
}

func (s *StubStore) Cancel() {
	return
}
