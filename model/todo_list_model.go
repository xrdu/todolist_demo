package model

import (
	"time"
)

type Todo struct {
	Id         string
	Title      string
	Body       string
	CreateUser string
	CreateTime time.Time
	UpdateTime time.Time
}

type TodoList struct {
	data []*Todo
}

func NewTodoList() *TodoList {
	return &TodoList{
		data: make([]*Todo, 0),
	}
}

func (t *TodoList) Get(id string) *Todo {
	for _, todo := range t.data {
		if todo.Id == id {
			return todo
		}
	}
	return nil
}

func (t *TodoList) Create(todo Todo) {
	t.data = append(t.data, &todo)
}

func (t *TodoList) ListAll() []*Todo {
	return t.data
}

func (t *TodoList) Delete(id string) {
	for i, todo := range t.data {
		if todo.Id == id {
			t.data[i] = nil
		}
	}
}

func (t *TodoList) Update(todo *Todo) {
	for i, v := range t.data {
		if v.Id == todo.Id {
			t.data[i] = todo
		}
	}
}

func (t *TodoList) ListByUser(user string) []*Todo {
	ret := make([]*Todo, 0)
	for _, todo := range t.data {
		if todo.CreateUser == user {
			ret = append(ret, todo)
		}
	}
	return ret
}
