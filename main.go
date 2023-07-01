package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"log"
)

func main() {
	
	router := chi.NewRouter()
	
	v1Router := GetV1Router()

	router.Mount("/v1", v1Router)

	log.Println("Server started on port 3000..")
	http.ListenAndServe(":3000", router)

}