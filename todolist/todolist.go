package todolist

import (
	"fmt"
)

type ToDoList struct {
	toDos []string
}

func New(things ...string) *ToDoList {
	todolist := ToDoList{}
	todolist.toDos = append(todolist.toDos, things...)

	return &todolist
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
