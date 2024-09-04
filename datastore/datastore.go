package datastore

import (
	"academy/todoapp/internal/model"
)

type TodoStore struct {
	UserTodos map[string]*model.TodoList
}

var todoList1 = model.NewTodoList("clean", "cook")
var todoList2 = model.NewTodoList("lunch", "train")

var Store = &TodoStore{
	UserTodos: map[string]*model.TodoList{
		"list1": todoList1,
		"list2": todoList2,
	},
}