package datastore

import (
	"academy/todoapp/todolist"
	"log"
	"fmt"
)

type TodoStore struct {
	userTodos map[string]todolist.TodoList
}

func userExists(datastore *TodoStore, listName string) bool {
	_, exist := datastore.userTodos[listName]
	return exist
}

func (datastore *TodoStore) CreateTodoList(listName string, todos []string) error {
	if userExists(datastore, listName) {
		return fmt.Errorf("listName - %s taken, please change", listName)
	}
	todoList := todolist.New(todos...)
	datastore.userTodos[listName] = *todoList
	log.Printf("Todo list - %s created successfully for user", listName)
	return nil
}

func (datastore *TodoStore) GetTodos(listName string) (string, error) {
	if !userExists(datastore, listName) {
		return "", fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, getting todos", listName)
	return todolist.Stringify(datastore.userTodos[listName]), nil
}

func (datastore *TodoStore) AddTodos(listName string, todos []string) error {
	if !userExists(datastore, listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, adding todos", listName)
	todoList := datastore.userTodos[listName]
	todoList.AddTodos(todos...)
	return nil
}

func (datastore *TodoStore) UpdateStatus(listName string, todo string, status string) error {
	if !userExists(datastore, listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, updating status", listName)
	todoList := datastore.userTodos[listName]
	return todoList.UpdateStatus(todo, status)
}

func (datastore *TodoStore) DeleteTodos(listName string, todos []string) error {
	if !userExists(datastore, listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, deleting todos", listName)
	todoList := datastore.userTodos[listName]
	return todoList.DeleteTodos(todos...)
}