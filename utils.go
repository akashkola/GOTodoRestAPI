package main

import (
	"encoding/json"
	"io"
	"strconv"
	"net/http"
	"github.com/go-chi/chi/v5"
	"errors"
)


func ConverTodoToJson(todo *TODO) ([]byte, error) {
	jsonTodo, err := json.Marshal(todo)
	return jsonTodo, err
}

func ConverTodosIntoJsonList(todos []*TODO) ([]byte, error) {
	jsonTodos, err := json.Marshal(todos)
	return jsonTodos, err
}

func ConvertBodyIntoTODO(body io.ReadCloser) (*TODO, error) {
	todo := &TODO{}
	err := json.NewDecoder(body).Decode(&todo)
	return todo, err
}

func RespondWithJson(data []byte, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func ConvertTODOIdIntoInt(todoId string) (int, error) {
	return strconv.Atoi(todoId)
}

func GetURLParam(r *http.Request, parameter string) (string) {
	return chi.URLParam(r, parameter)
}

func GetTODOFromTodos(t []*TODO,id int) *TODO {
	var todo *TODO = nil
	for _, v := range t {
		if id == v.Id {
			todo = v
			break
		}
	}
	return todo
}

func GetIndexFromTodos(id int, t []*TODO) (int, error) {
	for index, v := range t {
		if id == v.Id {
			return index, nil
		}
	}
	return -1, errors.New("Todo not found")
}

func RemoveElementFromTodos(t []*TODO, index int) []*TODO {
	t = append(t[:index], t[index+1:]...)
	return t
}