package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"academy/todoapp/internal/utils"
)

const (
	planned = "planned"
	done = "done"
)

type Todo struct {
	Description string
	Status string
}

type TodoList struct {
	ToDos []*Todo
	Note  string
}

func NewTodoList(descriptions ...string) *TodoList {
	todoList := TodoList{}
	todoList.ToDos = []*Todo{}
	todoList.AddTodos(descriptions...)

	if len(todoList.ToDos) == 0 {
		todoList.Note = "Nothing to do so far, but you can add some."
	}

	return &todoList
}

func todoExists(todos []*Todo, description string) (int, *Todo, bool) {
	for idx, todo := range todos {
		if todo.Description == description{
			return idx, todo, true
		}
	}
	return -1, nil, false
}

func (todoList *TodoList) AddTodos(descriptions ...string) {
	for _, description := range descriptions {
		todo := &Todo{
			Description: description,
			Status: planned,
		}
		todoList.ToDos = append(todoList.ToDos, todo)
	}
}

func (todoList *TodoList) UpdateStatus(description string, status string) error {
	_, todo, exists := todoExists(todoList.ToDos, description)
	if !exists {
		return fmt.Errorf("todo - %s does not exist", description)
	}
	todo.Status = status
	return nil
}

func (todoList *TodoList) DeleteTodos(descriptions ...string) error {
	var missingTodos string
	for _, description := range descriptions {
		idx, _, exists := todoExists(todoList.ToDos, description)
		if !exists {
			missingTodos = fmt.Sprintf("%s %s", description, missingTodos)
			continue
		}
		todoList.ToDos = utils.RemoveFromSliceByIndex(todoList.ToDos, idx)
	}
	if missingTodos == "" {
		return nil
	}
	return fmt.Errorf("todo - %sdo not exist", missingTodos)
}

func StringifyTodo(todoList TodoList) string {
	if len(todoList.ToDos) == 0 {
		return todoList.Note
	}

	var stringifiedTodos string
	i := 1 
	for _, todo := range todoList.ToDos {
		stringifiedTodos = fmt.Sprintf("%s%d. %s (%s) ", stringifiedTodos, i, todo.Description, todo.Status)
		i++
	}

	return stringifiedTodos
}

func GetJsonToDos(todolist TodoList) (string, error) {
	log.Print("Marshalling Todos...")
	jsonToDo, err := json.Marshal(todolist)

	if err != nil {
		log.Fatalf("Failed to marshal todos: %s", err.Error())
	}

	return string(jsonToDo), err
}

func WriteToJsonFile(todolist TodoList) error {
	file, err := os.Create("todos.json")

	if err != nil {
		log.Fatalf("Failed to create file: %s", err.Error())
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	err = encoder.Encode(todolist)

	if err != nil {
		log.Fatalf("Failed to write JSON to file: %s", err.Error())
	} else {
		log.Println("JSON data written to file successfully")
	}

	return err
}

func (todoList *TodoList) ReadFromJsonFile(name string) error {
	jsonData, err := os.ReadFile(name)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err.Error())
		return err
	}

	err = json.Unmarshal(jsonData, todoList)

	if err != nil {
		log.Fatalf("Failed to unmarshal json data: %s", err.Error())
	} else {
		log.Println("JSON data unmarshalled successfully")
	}

	return err
}
