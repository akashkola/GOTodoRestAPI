package main

import (
	"net/http"
	"log"
)

func HandlerUpdateStatus(w http.ResponseWriter, r *http.Request)  {
	
	// Get todoId from Url
	todoIdUrlParam := GetURLParam(r, "todoId")
	todoId, err := ConvertTODOIdIntoInt(todoIdUrlParam)
	if err != nil {
		log.Printf("error in conversion of \"%v\" to integer\n", todoIdUrlParam)
		w.WriteHeader(400)
		return
	}

	// Get Todo from Todos using todoId
	var todo *TODO
	todo = GetTODOFromTodos(Todos, todoId)
	if todo == nil {
		log.Printf("error todo with %v not found\n", todoId)
		w.WriteHeader(404)
		return
	}

	// Convert Body into TODO struct
	requestedTodo, err := ConvertBodyIntoTODO(r.Body)
	if err != nil {
		log.Println("error in parsing user requested data into todo")
		w.WriteHeader(400)
		return
	}

	if err != nil {
		log.Println("error title required")
		w.WriteHeader(400)
		return
	}

	// Update todo data
	requestedTodo.Title = todo.Title
	todo.Status = requestedTodo.Status
	requestedTodo.Id = todo.Id

	// Respond with added todo 
	todoTemp, err := ConverTodoToJson(requestedTodo)
	if err != nil {
		log.Println("error in converting todo to json")
		w.WriteHeader(500)
		return
	}
	RespondWithJson(todoTemp, w)
}