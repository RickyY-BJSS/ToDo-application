package repository_test

import (
	"context"
	"testing"

	"academy/todoapp/datastore"
	"academy/todoapp/internal/db"
	"academy/todoapp/internal/model"

	cmp "github.com/google/go-cmp/cmp"
)

var ctx context.Context = context.WithValue(context.TODO(), "fakeKey", "fakeID")

func TestRWMutexLock(t *testing.T) {
	t.Run("Test RWMutex on datastore", func(t *testing.T) {
		errorChannel := make(chan error)
		go func(errorChannel *chan error) {
			*errorChannel <- db.CreateTodos(ctx, "list3", []string{"yoga", "boxing"})
			
		}(&errorChannel)

		go func(errorChannel *chan error) {
			*errorChannel <- db.CreateTodos(ctx, "list3", []string{"party", "sleep"})
			
		}(&errorChannel)
		
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

func TestCreateTodos(t *testing.T) {
	t.Run("Successfully create todos", func(t *testing.T) {
		datastore := datastore.Store
		err := db.CreateTodos(ctx, "list3", []string{"yoga", "boxing"})

		if err != nil {
			t.Fatalf("Fail to create: %s", err.Error())
		}

		got, ok := datastore.UserTodos["list3"]
		want := model.TodoList{
			ToDos: []*model.Todo{
				{
					Description: "yoga",
					Status: "planned",
				},
				{
					Description: "boxing",
					Status: "planned",
				},
			},
			Note: "",
		}

		if !ok {
			t.Error("Todo not created")
		}

		if !cmp.Equal(*got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
