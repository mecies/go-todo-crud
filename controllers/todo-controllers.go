package controllers

import (
	"encoding/json"
	"fmt"
	"go-30/todo/config"
	"go-30/todo/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

type TodoRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoResponse struct {
	TodoRequest
	ID uint `json:"id"`
}

// Create todo data to database by run this function
func CreateTodo(responseWriter http.ResponseWriter, request *http.Request) {
	var data TodoRequest

	// Binding request body json to request body struct
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	// Matching todo models struct with todo request struct
	todo := models.Todo{}
	todo.ID = data.ID
	todo.Name = data.Name
	todo.Description = data.Description
	todo.Completed = false

	// Querying to database
	result := db.Create(&todo)

	if result.Error != nil {
		http.Error(responseWriter, "Couldn't create a todo", http.StatusBadRequest)
		return
	}

	// Matching result to create response
	var response TodoResponse

	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	json.NewEncoder(responseWriter).Encode(response)
}

func GetAllTodos(responseWriter http.ResponseWriter, request *http.Request) {

	var todos []models.Todo

	err := db.Find(&todos)

	if err.Error != nil {
		http.Error(responseWriter, "Couldn't get all todos", http.StatusBadRequest)
		return
	}

	json.NewEncoder(responseWriter).Encode(todos)
}

func UpdateTodo(responseWriter http.ResponseWriter, request *http.Request) {

	var data TodoRequest

	// Defining request parameter to get todo id
	vars := mux.Vars(request)
	todoId := vars["id"]

	// Binding request body json to request body struct
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	// Initiate models todo
	todo := models.Todo{}

	// Querying find todo data by todo id from request parameter
	todoById := db.Where("id = ?", todoId).First(&todo)

	if todoById.Error != nil {
		http.Error(responseWriter, "Couldn't find todo", http.StatusBadRequest)
		return
	}

	// Matching todo request with todo models
	todo.Name = data.Name
	todo.Description = data.Description
	todo.Completed = data.Completed

	// Update new todo data
	result := db.Save(&todo)

	if result.Error != nil {
		http.Error(responseWriter, "Couldn't update todo", http.StatusBadRequest)
		return
	}

	// Matching result to todo response struct
	var response TodoResponse

	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description
	response.Completed = todo.Completed

	json.NewEncoder(responseWriter).Encode(response)

}

func DeleteTodo(responseWriter http.ResponseWriter, request *http.Request) {
	todo := models.Todo{}

	// Defining request parameter to get todo id
	vars := mux.Vars(request)
	todoId := vars["id"]

	// Querying delete todo by id
	delete := db.Where("id = ?", todoId).Unscoped().Delete(&todo)
	fmt.Println(delete)

	fmt.Fprintf(responseWriter, "deleted %s", todoId)
}
