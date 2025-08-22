package models

import "github.com/google/uuid"

type Todo struct {
ID        string `json:"id"`
Title     string `json:"title"`
Completed bool   `json:"completed"`
}

func NewTodo(title string) Todo {
return Todo{ID: uuid.NewString(), Title: title, Completed: false}
}

