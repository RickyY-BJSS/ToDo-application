package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"academy/todoapp/internal/handler"
)

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
		handler.CreateTodos(recorder, fakeRequest)

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
		handler.CreateTodos(recorder, fakeRequest)

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
		handler.CreateTodos(recorder, fakeRequest)

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
		handler.CreateTodos(recorder, fakeRequest)

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
        		"drink",
        		"swim"
    		]
		}`
		fakeRequest := httptest.NewRequest("POST", "/todo/create", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		handler.CreateTodos(recorder, fakeRequest)

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
		handler.GetTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestDeleteTodos(t *testing.T) {
	requestBody := `{
		"descriptions": [
			"drink",
			"swim"
		]
	}`
	t.Run("List name missed, 400", func(t *testing.T) {
		fakeRequest := httptest.NewRequest("get", "/todo/listx/delete", strings.NewReader(requestBody))
		recorder := httptest.NewRecorder()
		handler.DeleteTodos(recorder, fakeRequest)

		got := recorder.Code
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
