package datastore

import (
	"academy/todoapp/todolist"
	"log"
	"fmt"
)

type TodoStore struct {
	UserTodos map[string]*todolist.TodoList
}

func listExists(todoStore *TodoStore, listName string) bool {
	_, exist := todoStore.UserTodos[listName]
	return exist
}

func (todoStore *TodoStore) CreateTodos(listName string, descriptions []string) error {
	if listExists(todoStore, listName) {
		return fmt.Errorf("listName - %s taken, please change", listName)
	}
	todoList := todolist.New(descriptions...)
	todoStore.UserTodos[listName] = todoList
	log.Printf("Todo list - %s created successfully for user", listName)
	return nil
}

func (todoStore *TodoStore) GetTodos(listName string) (string, error) {
	if !listExists(todoStore, listName) {
		return "", fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, getting todos", listName)
	return todolist.Stringify(*todoStore.UserTodos[listName]), nil
}

func (todoStore *TodoStore) AddTodos(listName string, descriptions []string) error {
	if !listExists(todoStore, listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, adding todos", listName)
	todoList := todoStore.UserTodos[listName]
	todoList.AddTodos(descriptions...)
	return nil
}

func (todoStore *TodoStore) UpdateStatus(listName string, description string, status string) error {
	if !listExists(todoStore, listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, updating status", listName)
	todoList := todoStore.UserTodos[listName]
	return todoList.UpdateStatus(description, status)
}

func (todoStore *TodoStore) DeleteTodos(listName string, descriptions []string) error {
	if !listExists(todoStore, listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	log.Printf("List - %s found, deleting todos", listName)
	todoList := todoStore.UserTodos[listName]
	return todoList.DeleteTodos(descriptions...)
}