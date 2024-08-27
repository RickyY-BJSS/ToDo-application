package todolist

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
)

type TodoList struct {
	ToDos []string
	Note  string
}

func New(things ...string) *TodoList {
	todoList := TodoList{}
	todoList.ToDos = append(todoList.ToDos, things...)

	if len(todoList.ToDos) == 0 {
		todoList.Note = "Nothing to do so far, but you can add some."
	}

	return &todoList
}

func String(todoList TodoList) string {
	if len(todoList.ToDos) == 0 {
		return todoList.Note
	}

	var stringifiedTodos string
	for i, thing := range todoList.ToDos {
		stringifiedTodos = fmt.Sprintf("%s%d. %s ", stringifiedTodos, i+1, thing)
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
