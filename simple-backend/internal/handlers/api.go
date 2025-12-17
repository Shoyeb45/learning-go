package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"simple-backend/internal/middleware"
)

func Handler(r *chi.Mux) {
	// global middleware
	r.Use(chimiddle.StripSlashes);


	// register /accounts route
	r.Route("/account", func (router chi.Router) {
		// add authorization middleware
		router.Use(middleware.Authorization);

		router.Get("/coins", GetCoinBalance);
	})
}