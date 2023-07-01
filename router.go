package main

import "github.com/go-chi/chi/v5"

func GetV1Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/todo", HandlerGetTODOs)
	router.Post("/todo", HandlerAddTODO)
	router.Get("/todo/{todoId}", HandlerGetSpecificTODO)
	router.Patch("/todo/{todoId}", HandlerUpdateTODO)
	router.Delete("/todo/{todoId}", HandlerDeleteTODO)

	return router
}