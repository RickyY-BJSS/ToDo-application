package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"academy/todoapp/datastore"
	"academy/todoapp/internal/db"
	"academy/todoapp/internal/handler"
	"academy/todoapp/internal/model"
	"academy/todoapp/internal/service"
	"academy/todoapp/internal/utils"
)


var todoStore *datastore.TodoStore
var traceIDKey utils.ContextKey

func main() {
	// intial data, dev env only
	todoList1 := model.NewTodoList("clean", "cook")
	todoList2 := model.NewTodoList("lunch", "train")
	
	todoStore = &datastore.TodoStore{
		UserTodos: map[string]*model.TodoList{
			"list1": todoList1,
			"list2": todoList2,
		},
	}

	repo := db.NewRepo(todoStore)
    todoService := service.New(repo)
    todoHandler := handler.New(todoService)

	router := mux.NewRouter()
	router.HandleFunc("/todo/create", todoHandler.CreateTodos).Methods(http.MethodPost)
	router.HandleFunc("/todo/{listName}", todoHandler.GetTodos).Methods(http.MethodGet)
	router.HandleFunc("/todo/{listName}/add", todoHandler.AddTodos).Methods(http.MethodPost)
	router.HandleFunc("/todo/{listName}/update-status", todoHandler.UpdateStatus).Methods(http.MethodPost)
	router.HandleFunc("/todo/{listName}/delete", todoHandler.DeleteTodos).Methods(http.MethodPost)
	router.Use(loggingMiddleware)

	log.Println("Starting App")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.RequestURI)

        traceID := r.Header.Get("X-Trace-ID")
		traceIDKey = "traceID"
        ctx := context.WithValue(r.Context(), traceIDKey, traceID)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}