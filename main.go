package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Todo represents a single todo item
// This is a struct - a way to group related data together
type Todo struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

// TodoList represents a collection of todos
// This is a slice of Todo structs
type TodoList []Todo

// Request/Response structures for API
type CreateTodoRequest struct {
	Task string `json:"task"`
}

type UpdateTodoRequest struct {
	Complete bool `json:"complete"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Global variable to store our todos
var todos TodoList

var nextID int = 1

// addTodo adds a new todo to the list
// This function demonstrates:
// - Function parameters and return values
// - Working with structs
// - Appending to slices
func addTodo(task string) Todo {
	newTodo := Todo{
		ID:       nextID,
		Task:     task,
		Complete: false,
	}

	todos = append(todos, newTodo)
	nextID++

	return newTodo
}

// getTodos returns all todos
func getTodos() TodoList {
	return todos
}

// getTodoByID finds a todo by ID
func getTodoByID(id int) (*Todo, bool) {
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], true
		}
	}
	return nil, false
}

// completeTodo marks a todo as complete
// This function demonstrates:
// - Loops with index
// - Working with pointers
// - Error handling
func completeTodo(id int) (*Todo, bool) {
	for i := range todos {
		if todos[i].ID == id {
			if todos[i].Complete {
				return &todos[i], false // Already complete
			}
			todos[i].Complete = true
			return &todos[i], true
		}
	}
	return nil, false // Not found
}

// deleteTodo removes a todo from the list
// This function demonstrates:
// - Slice manipulation
// - Working with indices
func deleteTodo(id int) (*Todo, bool) {
	for i, todo := range todos {
		if todo.ID == id {
			// Remove the todo by creating a new slice without it
			todos = append(todos[:i], todos[i+1:]...)
			return &todo, true
		}
	}
	return nil, false
}

// HTTP Handlers

// handleGetTodos returns all todos
func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	todos := getTodos()
	response := APIResponse{
		Success: true,
		Message: "Todos retrieved successfully",
		Data:    todos,
	}
	
	json.NewEncoder(w).Encode(response)
}

// handleCreateTodo creates a new todo
func handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := APIResponse{
			Success: false,
			Message: "Invalid JSON format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if strings.TrimSpace(req.Task) == "" {
		response := APIResponse{
			Success: false,
			Message: "Task cannot be empty",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	todo := addTodo(req.Task)
	response := APIResponse{
		Success: true,
		Message: "Todo created successfully",
		Data:    todo,
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// handleGetTodo returns a specific todo by ID
func handleGetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(path)
	if err != nil {
		response := APIResponse{
			Success: false,
			Message: "Invalid todo ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	todo, found := getTodoByID(id)
	if !found {
		response := APIResponse{
			Success: false,
			Message: "Todo not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	response := APIResponse{
		Success: true,
		Message: "Todo retrieved successfully",
		Data:    todo,
	}
	
	json.NewEncoder(w).Encode(response)
}

// handleUpdateTodo updates a todo (mark as complete/incomplete)
func handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(path)
	if err != nil {
		response := APIResponse{
			Success: false,
			Message: "Invalid todo ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	var req UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := APIResponse{
			Success: false,
			Message: "Invalid JSON format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	todo, found := getTodoByID(id)
	if !found {
		response := APIResponse{
			Success: false,
			Message: "Todo not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	todo.Complete = req.Complete
	
	response := APIResponse{
		Success: true,
		Message: "Todo updated successfully",
		Data:    todo,
	}
	
	json.NewEncoder(w).Encode(response)
}

// handleDeleteTodo deletes a todo
func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(path)
	if err != nil {
		response := APIResponse{
			Success: false,
			Message: "Invalid todo ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	todo, found := deleteTodo(id)
	if !found {
		response := APIResponse{
			Success: false,
			Message: "Todo not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	response := APIResponse{
		Success: true,
		Message: "Todo deleted successfully",
		Data:    todo,
	}
	
	json.NewEncoder(w).Encode(response)
}

// handleCompleteTodo marks a todo as complete (convenience endpoint)
func handleCompleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/todos/")
	path = strings.TrimSuffix(path, "/complete")
	id, err := strconv.Atoi(path)
	if err != nil {
		response := APIResponse{
			Success: false,
			Message: "Invalid todo ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	todo, success := completeTodo(id)
	if !success {
		if todo == nil {
			response := APIResponse{
				Success: false,
				Message: "Todo not found",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
		} else {
			response := APIResponse{
				Success: false,
				Message: "Todo is already complete",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		}
		return
	}
	
	response := APIResponse{
		Success: true,
		Message: "Todo marked as complete",
		Data:    todo,
	}
	
	json.NewEncoder(w).Encode(response)
}

// handleHealth provides a health check endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := APIResponse{
		Success: true,
		Message: "Todo API is running",
		Data: map[string]interface{}{
			"total_todos": len(todos),
		},
	}
	
	json.NewEncoder(w).Encode(response)
}

// main is the entry point of the program
// This function demonstrates:
// - HTTP server setup
// - Route handling
// - Web API structure
func main() {
	fmt.Println("�� Starting Todo API Server!")
	fmt.Println("A beginner-friendly Go web API")
	
	// Set up routes
	http.HandleFunc("/", handleHealth)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetTodos(w, r)
		case http.MethodPost:
			handleCreateTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/complete") {
			handleCompleteTodo(w, r)
		} else {
			switch r.Method {
			case http.MethodGet:
				handleGetTodo(w, r)
			case http.MethodPut:
				handleUpdateTodo(w, r)
			case http.MethodDelete:
				handleDeleteTodo(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		}
	})
	
	// Add some sample todos
	addTodo("Learn Go programming")
	addTodo("Build a web API")
	addTodo("Test with Postman")
	
	fmt.Println("📋 Sample todos added!")
	fmt.Println("🌐 Server starting on http://localhost:8080")
	fmt.Println("\n📖 Available endpoints:")
	fmt.Println("  GET    /health              - Health check")
	fmt.Println("  GET    /todos               - Get all todos")
	fmt.Println("  POST   /todos               - Create a new todo")
	fmt.Println("  GET    /todos/{id}          - Get a specific todo")
	fmt.Println("  PUT    /todos/{id}          - Update a todo")
	fmt.Println("  DELETE /todos/{id}          - Delete a todo")
	fmt.Println("  POST   /todos/{id}/complete - Mark todo as complete")
	fmt.Println("\n🚀 Server is running! Use Postman to test the endpoints.")
	
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
