package main

import (
	"net/http"
	"log"
)

func HandlerGetTODOs(w http.ResponseWriter, r *http.Request)  {

	// Convert Todos into json
	todosJson, err := ConverTodosIntoJsonList(Todos)
	if err != nil {
		log.Println("error occured in conversion of todos to json")
		w.WriteHeader(500)
		return
	}

	// Respond with todos json
	RespondWithJson(todosJson, w)
}