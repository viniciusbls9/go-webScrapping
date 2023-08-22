package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func handlerRouters() (router *chi.Mux) {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/scrapper/{productName}", scrapper)

	r.Mount("/v1", v1Router)

	return router
}