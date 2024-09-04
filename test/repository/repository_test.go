package repository_test

import (
	"os"
	"testing"

	"academy/todoapp/datastore"
	"academy/todoapp/internal/db"
	"academy/todoapp/internal/model"
)

var todoStore *datastore.TodoStore
var repo *db.Repository

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

    os.Exit(code)
}

func setup() {
	todoList1 := model.NewTodoList("clean", "cook")
	todoList2 := model.NewTodoList("lunch", "train")
	
	todoStore = &datastore.TodoStore{
		UserTodos: map[string]*model.TodoList{
			"list1": todoList1,
			"list2": todoList2,
		},
	}

	repo = db.NewRepo(todoStore) 
}

func TestRWMutexLock(t *testing.T) {
	t.Run("Test RWMutex on datastore", func(t *testing.T) {
		errorChannel := make(chan error)
		go func(errorChannel *chan error) {
			*errorChannel <- repo.CreateTodos("list3", []string{"yoga", "boxing"})
			
		}(&errorChannel)

		go func(errorChannel *chan error) {
			*errorChannel <- repo.CreateTodos("list3", []string{"party", "sleep"})
			
		}(&errorChannel)

		defer setup()
		
		err1 := <- errorChannel
		err2 := <- errorChannel

		if err1 != nil {
			t.Fatal("list3 should be created successfully")
		}

		got := err2.Error()
		want := "listName - list3 taken, please change" 
		if got != want  {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
