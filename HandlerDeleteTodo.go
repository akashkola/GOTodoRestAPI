package main

import (
	"net/http"
	"log"
)

func HandlerDeleteTODO(w http.ResponseWriter, r *http.Request)  {

	// Get todoId from Url
	todoId, err := ConvertTODOIdIntoInt(GetURLParam(r, "todoId"))
	if err != nil {
		log.Println("error in conversion of todoId to integer")
		w.WriteHeader(400)
		return
	}

	// Remove Element from Todos
	index, err := GetIndexFromTodos(todoId, Todos)
	if err != nil {
		log.Printf("todo element with %v not found", todoId)
		w.WriteHeader(404)
		return
	}
	Todos = RemoveElementFromTodos(Todos, index)
}