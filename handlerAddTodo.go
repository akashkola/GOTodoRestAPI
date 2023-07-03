package main

import (
	"net/http"
	"log"
)

func HandlerAddTODO(w http.ResponseWriter, r *http.Request)  {
	
	// Convert Body into TODO struct
	todo, err := ConvertBodyIntoTODO(r.Body)
	if err != nil {
		log.Println("error in parsing user requested data into todo")
		w.WriteHeader(400)
		return
	}

	err = ValidateTodo(todo)
	if err != nil {
		log.Println("error title required")
		w.WriteHeader(400)
		return
	}

	// Add todo to Todos
	TODOsId += 1
	todo.Id = TODOsId
	todo.Status = false
	Todos = append(Todos, todo)

	// Respond with added todo 
	todoTemp, err := ConverTodoToJson(todo)
	if err != nil {
		log.Println("error in converting todo to json")
		w.WriteHeader(500)
		return
	}
	RespondWithJson(todoTemp, w)

}