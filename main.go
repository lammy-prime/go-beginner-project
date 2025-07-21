package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Todo represents a single todo item
// This is a struct - a way to group related data together
type Todo struct {
	ID       int
	Task     string
	Complete bool
}

// TodoList represents a collection of todos
// This is a slice of Todo structs
type TodoList []Todo

// Global variable to store our todos
var todos TodoList
var nextID int = 1

// addTodo adds a new todo to the list
// This function demonstrates:
// - Function parameters and return values
// - Working with structs
// - Appending to slices
func addTodo(task string) {
	newTodo := Todo{
		ID:       nextID,
		Task:     task,
		Complete: false,
	}

	todos = append(todos, newTodo)
	nextID++

	fmt.Printf("✅ Added todo: %s (ID: %d)\n", task, newTodo.ID)
}

// listTodos displays all todos
// This function demonstrates:
// - Loops (for range)
// - Conditional statements (if/else)
// - String formatting
func listTodos() {
	if len(todos) == 0 {
		fmt.Println("📝 No todos found. Add some todos first!")
		return
	}

	fmt.Println("\n📋 Your Todo List:")
	fmt.Println("==================")

	for _, todo := range todos {
		status := "❌"
		if todo.Complete {
			status = "✅"
		}
		fmt.Printf("%s [%d] %s\n", status, todo.ID, todo.Task)
	}
	fmt.Println()
}

// completeTodo marks a todo as complete
// This function demonstrates:
// - Loops with index
// - Working with pointers
// - Error handling
func completeTodo(id int) {
	for i := range todos {
		if todos[i].ID == id {
			if todos[i].Complete {
				fmt.Printf("❌ Todo %d is already complete!\n", id)
			} else {
				todos[i].Complete = true
				fmt.Printf("✅ Marked todo %d as complete: %s\n", id, todos[i].Task)
			}
			return
		}
	}
	fmt.Printf("❌ Todo with ID %d not found!\n", id)
}

// deleteTodo removes a todo from the list
// This function demonstrates:
// - Slice manipulation
// - Working with indices
func deleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			// Remove the todo by creating a new slice without it
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Printf("🗑️  Deleted todo %d: %s\n", id, todo.Task)
			return
		}
	}
	fmt.Printf("❌ Todo with ID %d not found!\n", id)
}

// showMenu displays the available commands
func showMenu() {
	fmt.Println("\n🎯 Todo App - Available Commands:")
	fmt.Println("1. add <task>     - Add a new todo")
	fmt.Println("2. list           - Show all todos")
	fmt.Println("3. complete <id>  - Mark a todo as complete")
	fmt.Println("4. delete <id>    - Delete a todo")
	fmt.Println("5. help           - Show this menu")
	fmt.Println("6. quit           - Exit the application")
	fmt.Println()
}

// getInput reads user input from the console
// This function demonstrates:
// - Working with the bufio package
// - String manipulation
// - Error handling
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("❌ Error reading input:", err)
		return ""
	}

	// Remove the newline character from the end
	return strings.TrimSpace(input)
}

// processCommand handles user commands
// This function demonstrates:
// - String splitting
// - Switch statements
// - Type conversion (strconv)
func processCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := strings.ToLower(parts[0])

	switch command {
	case "add":
		if len(parts) < 2 {
			fmt.Println("❌ Please provide a task description!")
			return
		}
		task := strings.Join(parts[1:], " ")
		addTodo(task)

	case "list":
		listTodos()

	case "complete":
		if len(parts) < 2 {
			fmt.Println("❌ Please provide a todo ID!")
			return
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("❌ Invalid ID! Please enter a number.")
			return
		}
		completeTodo(id)

	case "delete":
		if len(parts) < 2 {
			fmt.Println("❌ Please provide a todo ID!")
			return
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("❌ Invalid ID! Please enter a number.")
			return
		}
		deleteTodo(id)

	case "help":
		showMenu()

	case "quit":
		fmt.Println("👋 Thanks for using Todo App! Goodbye!")
		os.Exit(0)

	default:
		fmt.Printf("❌ Unknown command: %s\n", command)
		fmt.Println("Type 'help' to see available commands.")
	}
}

// main is the entry point of the program
// This function demonstrates:
// - Program structure
// - Infinite loops
// - User interaction
func main() {
	fmt.Println("🚀 Welcome to Todo App!")
	fmt.Println("A beginner-friendly Go application")

	// Show the menu initially
	showMenu()

	// Main program loop
	for {
		fmt.Print("📝 Enter command: ")
		input := getInput()
		processCommand(input)
	}
}
