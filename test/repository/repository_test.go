package repository_test

import (
	"testing"
	"context"

	"academy/todoapp/internal/db"
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
