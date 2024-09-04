package handler_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"academy/todoapp/datastore"
	"academy/todoapp/internal/db"
	"academy/todoapp/internal/handler"
	"academy/todoapp/internal/model"
	"academy/todoapp/internal/service"
)

var todoHandler *handler.TodoHandler

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

    os.Exit(code)
}

func setup() {
	todoList1 := model.NewTodoList("clean", "cook")
	todoList2 := model.NewTodoList("lunch", "train")
	
	todoStore := &datastore.TodoStore{
		UserTodos: map[string]*model.TodoList{
			"list1": todoList1,
			"list2": todoList2,
		},
	}

	repo := db.NewRepo(todoStore)
    todoService := service.New(repo)
    todoHandler = handler.New(todoService)
}

func TestCreateTodos(t *testing.T) {
	t.Run("Successfully create todos, 202", func(t *testing.T) {
		requestBody := `{
			"listName": "listx",
			"descriptions": [
        		"drink",
        		"swim"
    		]
		}`
		fakeRequest := httptest.NewRequest("POST", "/todo/create", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		todoHandler.CreateTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusAccepted

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Get a bad request with fields that can't be recognised, 400", func(t *testing.T) {
		requestBody := `{
			"listName": "listx",
			"wrongfield": [
        		"drink",
        		"swim"
    		]
		}`
		fakeRequest := httptest.NewRequest("POST", "/todo/create", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		todoHandler.CreateTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Get a bad request with invalid JSON format, 400", func(t *testing.T) {
		requestBody := `{
			"listName": "listx",
			descriptions: [
        		drink,
        		"swim"
    		]
		}`
		fakeRequest := httptest.NewRequest("POST", "/todo/create", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		todoHandler.CreateTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Empty list name and todos, 400", func(t *testing.T) {
		requestBody := `{
			"wrongname": "listx",
			"wronddescription": [
        		drink,
        		"swim"
    		]
		}`
		fakeRequest := httptest.NewRequest("POST", "/todo/create", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		todoHandler.CreateTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("The list name is taken, 500", func(t *testing.T) {
		requestBody := `{
			"listName": "list1",
			"descriptions": [
        		drink,
        		"swim"
    		]
		}`
		fakeRequest := httptest.NewRequest("POST", "/todo/create", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		todoHandler.CreateTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusInternalServerError

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})	
}

func TestGetTodos(t *testing.T) {
	t.Run("List name doesn't exist, can't find todos, 404", func(t *testing.T) {
		requestBody := `{
			"listName": "list1",
			"descriptions": [
        		drink,
        		"swim"
    		]
		}`
		fakeRequest := httptest.NewRequest("get", "/todo/unknown", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		todoHandler.GetTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestDeleteTodos(t *testing.T) {
	t.Run("List name missed, 400", func(t *testing.T) {
		fakeRequest := httptest.NewRequest("get", "/todo/unknown", nil)
		recorder := httptest.NewRecorder()
		todoHandler.DeleteTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
