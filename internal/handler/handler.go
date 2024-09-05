package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"academy/todoapp/internal/model"
	"academy/todoapp/internal/service"
	"academy/todoapp/internal/utils"

	"github.com/gorilla/mux"
)


func CreateTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request model.TodosModificationRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid JSON: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	listName := request.ListName
	descriptions := request.Descriptions
	if listName == nil && descriptions == nil {
		http.Error(w, "Empty list name and todos", http.StatusBadRequest)
		return
	}

	err = service.CreateTodos(ctx, request.ListName, request.Descriptions)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	listName := mux.Vars(r)["listName"]
	todos, err := service.GetTodos(ctx, listName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	msg := fmt.Sprintf("Todos in List - %s fetched successfully", listName)
	utils.LogWithTraceID(ctx, msg)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func AddTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request model.TodosModificationRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid JSON: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	listName := request.ListName
	descriptions := request.Descriptions
	if listName == nil {
		http.Error(w, "List name missed", http.StatusBadRequest)
		return
	}

	if descriptions == nil {
		http.Error(w, "Cannot add nothing", http.StatusBadRequest)
		return
	}

	err = service.AddTodos(ctx, listName, descriptions)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	listName := mux.Vars(r)["listName"]
	queryParam := r.URL.Query()
	description := queryParam.Get("description")
	status := queryParam.Get("status")

	err := service.UpdateStatus(ctx, listName, description, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Status updated successfully in List - %s", listName)
	utils.LogWithTraceID(ctx, msg)

	w.WriteHeader(http.StatusAccepted)
}

func DeleteTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request model.TodosModificationRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid JSON: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	listName := request.ListName
	descriptions := request.Descriptions
	if listName == nil {
		http.Error(w, "List name missed", http.StatusBadRequest)
		return
	}

	if descriptions == nil {
		http.Error(w, "Cannot delete nothing", http.StatusBadRequest)
		return
	}

	err = service.DeleteTodos(ctx, listName, descriptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Todos deleted from List - %s successfully", *listName)
	utils.LogWithTraceID(ctx, msg)

	w.WriteHeader(http.StatusAccepted)
}
