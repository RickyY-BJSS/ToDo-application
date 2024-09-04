package datastore

import (
	"academy/todoapp/internal/model"
)

type TodoStore struct {
	UserTodos map[string]*model.TodoList
}
