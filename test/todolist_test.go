package test

import (
	"testing"
	"academy/todoapp/todolist"
)

func TestPrintToDos(t *testing.T) {
	t.Run("Print out all todos, should return e.g. 1. first_thing", func(t *testing.T) {
		toDoList := todolist.New("Code", "Cook")

		got := todolist.PrintToDos(*toDoList)
		want := "1. Code\n2. Cook\n"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	
	t.Run("0 todo in the list, should return empty string", func(t *testing.T) {
		toDoList := todolist.New()

		got := todolist.PrintToDos(*toDoList)
		want := ""
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}