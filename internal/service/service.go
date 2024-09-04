package service

import (
    "context"

    "academy/todoapp/internal/model"
	"academy/todoapp/internal/db"
)

func CreateTodos(ctx context.Context, listName *string, descriptions *[]string) error {
    var realName string
    var realDescriptions []string
    if listName == nil {
        realName = "Unnamed"
    } else {
        realName = *listName
    }

    if descriptions == nil {
        realDescriptions = []string{}
    } else {
        realDescriptions = *descriptions
    }
    return db.CreateTodos(ctx, realName, realDescriptions)
}

func GetTodos(ctx context.Context, listName string) (*model.TodoList, error) {
    return db.GetTodos(ctx, listName)
}

func AddTodos(ctx context.Context, listName *string, descriptions *[]string) error {
    return db.AddTodos(ctx, *listName, *descriptions)
}

func UpdateStatus(ctx context.Context, listName string, description string, status string) error {
    return db.UpdateStatus(ctx, listName, description, status)
}

func DeleteTodos(ctx context.Context, listName *string, descriptions *[]string)  error {
    return db.DeleteTodos(ctx, *listName, *descriptions)
}
