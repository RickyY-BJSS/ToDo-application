package todolist

import (
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

func PrintToDos(todolist ToDoList) string {
	var output string
	if len(todolist.toDos) == 0 {
		return "Nothing to do so far, but you can add some."
	}
	for i, thing := range todolist.toDos {
		output = fmt.Sprintf("%s%d. %s\n", output, i+1, thing)
	}
	return output
}
