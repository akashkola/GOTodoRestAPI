package main

type TODO struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}