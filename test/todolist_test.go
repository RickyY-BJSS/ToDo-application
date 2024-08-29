package test

import (
	"academy/todoapp/todolist"
	"os"
	"testing"
	"reflect"
)

func TestPrintToDos(t *testing.T) {
	t.Run("Print out all todos with empty note", func(t *testing.T) {
		todos := todolist.New("Code", "Cook")

		got, err := todolist.GetJsonToDos(*todos)

		if err != nil {
			t.Fatalf("Failed to parse todos: %s", err.Error())
		}

		want := "{\"ToDos\":{\"Code\":\"planned\",\"Cook\":\"planned\"},\"Note\":\"\"}"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	
	t.Run("0 todo in the list, should return with a note", func(t *testing.T) {
		todos := todolist.New()

		got, err := todolist.GetJsonToDos(*todos)

		if err != nil {
			t.Fatalf("Failed to parse todos: %s", err.Error())
		}

		want := "{\"ToDos\":{},\"Note\":\"Nothing to do so far, but you can add some.\"}"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Write todos to a Json format file", func(t *testing.T) {
		todos := todolist.New("Code", "Cook")

		if err := todolist.WriteToJsonFile(*todos); err != nil {
			t.Fatalf("Failed to output json data: %s", err.Error())
		}

		jsonTodos, err := os.ReadFile("todos.json")

		if err != nil {
			t.Fatalf("Failed to read file: %s", err.Error())
		}

		defer os.Remove("todos.json")

		got, err := CompactJson(jsonTodos)

		if err != nil {
			t.Fatalf("Failed to compact json data %s", err.Error())
		}

		want := "{\"ToDos\":{\"Code\":\"planned\",\"Cook\":\"planned\"},\"Note\":\"\"}"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Read todos to a Json format file", func(t *testing.T) {
		want := todolist.New("Code", "Cook")

		if err := todolist.WriteToJsonFile(*want); err != nil {
			t.Fatalf("Failed to output json data: %s", err.Error())
		}
		
		got := todolist.New()

		if err := got.ReadFromJsonFile("todos.json"); err != nil {
			t.Fatalf("Failed to read json data: %s", err.Error())
		}

		defer os.Remove("todos.json")

		if !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}