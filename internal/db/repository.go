package db

import (
	// "database/sql"
	"fmt"
	"sync"
	"context"

	"academy/todoapp/datastore"
	"academy/todoapp/internal/model"
	"academy/todoapp/internal/utils"
)

type Repository struct {
	//DB *sql.DB
	TodoStore *datastore.TodoStore
	RwMutex   *sync.RWMutex
}

var repo Repository = Repository{
	//DB: db,
	TodoStore: datastore.Store,
	RwMutex:   &sync.RWMutex{},
}

func listExists(listName string) bool {
	_, exist := repo.TodoStore.UserTodos[listName]
	return exist
}

func CreateTodos(ctx context.Context, listName string, descriptions []string) error {
	repo.RwMutex.Lock()
	defer repo.RwMutex.Unlock()

	if listExists(listName) {
		return fmt.Errorf("listName - %s taken, please change", listName)
	}
	todoList := model.NewTodoList(descriptions...)
	repo.TodoStore.UserTodos[listName] = todoList
	msg := fmt.Sprintf("Todo list - %s created successfully for user", listName)
	utils.LogWithTraceID(ctx, msg)
	return nil
}

func GetTodos(ctx context.Context, listName string) (*model.TodoList, error) {
	repo.RwMutex.RLock()
	defer repo.RwMutex.RUnlock()

	if !listExists(listName) {
		return nil, fmt.Errorf("listName - %s does not exist", listName)
	}
	msg := fmt.Sprintf("List - %s found, getting todos", listName)
	utils.LogWithTraceID(ctx, msg)
	return repo.TodoStore.UserTodos[listName], nil
}

func AddTodos(ctx context.Context, listName string, descriptions []string) error {
	repo.RwMutex.Lock()
	defer repo.RwMutex.Unlock()

	if !listExists(listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	msg := fmt.Sprintf("List - %s found, adding todos", listName)
	utils.LogWithTraceID(ctx, msg)
	todoList := repo.TodoStore.UserTodos[listName]
	todoList.AddTodos(descriptions...)
	msg = fmt.Sprintf("Todos added to List - %s successfully", listName)
	utils.LogWithTraceID(ctx, msg)
	return nil
}

func UpdateStatus(ctx context.Context, listName string, description string, status string) error {
	repo.RwMutex.Lock()
	defer repo.RwMutex.Unlock()

	if !listExists(listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	msg := fmt.Sprintf("List - %s found, updating status", listName)
	utils.LogWithTraceID(ctx, msg)
	todoList := repo.TodoStore.UserTodos[listName]
	return todoList.UpdateStatus(description, status)
}

func DeleteTodos(ctx context.Context, listName string, descriptions []string) error {
	repo.RwMutex.Lock()
	defer repo.RwMutex.Unlock()

	if !listExists(listName) {
		return fmt.Errorf("listName - %s does not exist", listName)
	}
	msg := fmt.Sprintf("List - %s found, deleting todos", listName)
	utils.LogWithTraceID(ctx, msg)
	todoList := repo.TodoStore.UserTodos[listName]
	return todoList.DeleteTodos(descriptions...)
}
