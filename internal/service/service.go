package service

import (
    "context"

    "academy/todoapp/internal/model"
	"academy/todoapp/internal/db"
)

// This service layer is redundant here, 
// as there's no additional business logic 
// other than the basic CRUD 

type TodoService struct {
    Repo *db.Repository
}

func New(repo *db.Repository) *TodoService {
    return &TodoService{Repo: repo}
}

func (s *TodoService) CreateTodos(ctx context.Context, listName *string, descriptions *[]string) error {
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
    return s.Repo.CreateTodos(ctx, realName, realDescriptions)
}

func (s *TodoService) GetTodos(ctx context.Context, listName string) (*model.TodoList, error) {
    return s.Repo.GetTodos(ctx, listName)
}

func (s *TodoService) AddTodos(ctx context.Context, listName *string, descriptions *[]string) error {
    return s.Repo.AddTodos(ctx, *listName, *descriptions)
}

func (s *TodoService) UpdateStatus(ctx context.Context, listName string, description string, status string) error {
    return s.Repo.UpdateStatus(ctx, listName, description, status)
}

func (s *TodoService) DeleteTodos(ctx context.Context, listName *string, descriptions *[]string)  error {
    return s.Repo.DeleteTodos(ctx, *listName, *descriptions)
}
