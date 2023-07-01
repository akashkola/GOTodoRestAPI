package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"log"
	"github.com/go-chi/cors"
)

func main() {
	
	router := chi.NewRouter()
	
	v1Router := GetV1Router()

	 router.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))


	router.Mount("/v1", v1Router)

	log.Println("Server started on port 3000..")
	http.ListenAndServe(":3000", router)

}
