package test

import (
	"academy/todoapp/datastore"
	"academy/todoapp/todolist"
	"os"
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

var todoStore *datastore.TodoStore


func TestMain(m *testing.M) {
	setup()

	code := m.Run()

    os.Exit(code)
}

func setup() {
	todoList1 := todolist.New("clean", "cook")
	todoList2 := todolist.New("lunch", "train")
	
	todoStore = &datastore.TodoStore{
		UserTodos: map[string]*todolist.TodoList{
			"list1": todoList1,
			"list2": todoList2,
		},
	}
}

func TestCreateTodos(t *testing.T) {
	t.Run("Successfully create a todo list", func(t *testing.T) {
		err := todoStore.CreateTodos("list3", []string{"yoga", "boxing"})

		defer setup()
		
		if err != nil {
			t.Fatalf("Failed to create todo list: %s", err.Error())
		}

		got := todoStore.UserTodos["list3"]

		want := todolist.New("yoga", "boxing")

		if !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Fail to create a todo list, duplicate list name", func(t *testing.T) {
		err := todoStore.CreateTodos("list2", []string{"yoga", "boxing"})

		got := err.Error()

		want := "listName - list2 taken, please change"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestGetTodos(t *testing.T) {
	t.Run("Successfully get a todo list", func(t *testing.T) {
		stringifiedTodos, err := todoStore.GetTodos("list2")

		if err != nil {
			t.Fatalf("Failed to get todo list: %s", err.Error())
		}

		got := stringifiedTodos

		want := "1. lunch (planned) 2. train (planned) "

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Fail to get a todo list, list name doesn't exist", func(t *testing.T) {
		stringifiedTodos, err := todoStore.GetTodos("list3")

		got := stringifiedTodos

		want := ""

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		got = err.Error()

		want = "listName - list3 does not exist"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestAddTodos(t *testing.T) {
	t.Run("Successfully add new todos to a list", func(t *testing.T) {
		err := todoStore.AddTodos("list2", []string{"dinner"})

		defer setup()

		if err != nil {
			t.Fatalf("Failed to add new todos to todo list: %s", err.Error())
		}

		got := todoStore.UserTodos["list2"]

		want := todolist.New("lunch", "train", "dinner")

		if !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Fail to add new todos, list name doesn't exist", func(t *testing.T) {
		err := todoStore.AddTodos("list3", []string{"dinner"})

		got := err.Error()

		want := "listName - list3 does not exist"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Successfully update the status of a todo", func(t *testing.T) {
		err := todoStore.UpdateStatus("list2", "lunch", "done")

		defer setup()

		if err != nil {
			t.Fatalf("Failed to add new todos to todo list: %s", err.Error())
		}

		got := todoStore.UserTodos["list2"].ToDos[0].Status

		want := "done"

		if !cmp.Equal(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Fail to update status, list name doesn't exist", func(t *testing.T) {
		err := todoStore.UpdateStatus("list3", "lunch", "done")

		got := err.Error()

		want := "listName - list3 does not exist"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Fail to update status, todo doesn't exist", func(t *testing.T) {
		err := todoStore.UpdateStatus("list2", "ride", "done")

		got := err.Error()

		want := "todo - ride does not exist"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestDeleteTodos(t *testing.T) {
	t.Run("Successfully delete todos", func(t *testing.T) {
		err := todoStore.DeleteTodos("list2", []string{"lunch"})

		defer setup()

		if err != nil {
			t.Fatalf("Failed to delete todos: %s", err.Error())
		}

		got := todoStore.UserTodos["list2"]

		want := todolist.New("train")

		if !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Fail to delete todos, list name doesn't exist", func(t *testing.T) {
		err := todoStore.DeleteTodos("list3", []string{"lunch"})

		got := err.Error()

		want := "listName - list3 does not exist"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Fail to delete, some todos don't exist", func(t *testing.T) {
		err := todoStore.DeleteTodos("list2", []string{"lunch", "ride"})

		defer setup()

		got := err.Error()

		want := "todo - ride do not exist"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		got2 := todoStore.UserTodos["list2"]

		want2 := todolist.New("train")

		if !cmp.Equal(got2, want2) {
			t.Errorf("got %v, want %v", got2, want2)
		}
	})
}
