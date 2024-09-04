package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"academy/todoapp/datastore"
	"academy/todoapp/internal/handler"
	"academy/todoapp/internal/utils"
)


var todoStore *datastore.TodoStore
var traceIDKey utils.ContextKey

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todo/create", handler.CreateTodos).Methods(http.MethodPost)
	router.HandleFunc("/todo/{listName}", handler.GetTodos).Methods(http.MethodGet)
	router.HandleFunc("/todo/{listName}/add", handler.AddTodos).Methods(http.MethodPost)
	router.HandleFunc("/todo/{listName}/update-status", handler.UpdateStatus).Methods(http.MethodPost)
	router.HandleFunc("/todo/{listName}/delete", handler.DeleteTodos).Methods(http.MethodPost)
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