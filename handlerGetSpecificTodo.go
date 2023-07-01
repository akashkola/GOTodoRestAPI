package main

import (
	"net/http"
	"log"
)

func HandlerGetSpecificTODO(w http.ResponseWriter, r *http.Request)  {

	// Get todoId from Url
	todoIdUrlParam := GetURLParam(r, "todoId")
	todoId, err := ConvertTODOIdIntoInt(todoIdUrlParam)
	if err != nil {
		log.Printf("error in conversion of \"%v\" to integer\n", todoIdUrlParam)
		w.WriteHeader(400)
		return
	}

	// Get Todo from Todos using todoId
	var todo *TODO = nil
	todo = GetTODOFromTodos(Todos, todoId)
	if todo == nil {
		log.Printf("error todo with %v not found\n", todoId)
		w.WriteHeader(404)
		return
	}

	// Respond with json format of todo
	todoTemp, err := ConverTodoToJson(todo)
	if err != nil {
		log.Println("error in conversion of todo to json")
		w.WriteHeader(500)
		return
	}
	RespondWithJson(todoTemp, w)

}