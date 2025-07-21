# 🚀 Beginner-Friendly Go Todo App

Welcome to your first Go project! This is a simple command-line todo application designed to teach you the fundamentals of Go programming.

## 📚 What You'll Learn

This project covers essential Go concepts:

- **Variables and Data Types**: `int`, `string`, `bool`
- **Structs**: Creating custom data types
- **Slices**: Dynamic arrays in Go
- **Functions**: Defining and calling functions
- **Control Flow**: `if/else`, `for` loops, `switch` statements
- **Error Handling**: Basic error checking
- **Input/Output**: Reading from console, printing to screen
- **String Manipulation**: Working with strings
- **Type Conversion**: Converting between types

## 🛠️ Prerequisites

Before you start, make sure you have Go installed on your computer:

1. **Install Go**: Download from [golang.org](https://golang.org/dl/)
2. **Verify Installation**: Open terminal and run:
   ```bash
   go version
   ```

## 🏃‍♂️ How to Run

1. **Navigate to the project directory**:

   ```bash
   cd go-beginner-project
   ```

2. **Run the application**:

   ```bash
   go run main.go
   ```

3. **Or build and run the executable**:
   ```bash
   go build main.go
   ./todo-app
   ```

## 🎯 How to Use the App

Once the app is running, you can use these commands:

- `add Buy groceries` - Add a new todo
- `list` - Show all todos
- `complete 1` - Mark todo with ID 1 as complete
- `delete 1` - Delete todo with ID 1
- `help` - Show available commands
- `quit` - Exit the application

### Example Session:

```
🚀 Welcome to Todo App!
A beginner-friendly Go application

🎯 Todo App - Available Commands:
1. add <task>     - Add a new todo
2. list           - Show all todos
3. complete <id>  - Mark a todo as complete
4. delete <id>    - Delete a todo
5. help           - Show this menu
6. quit           - Exit the application

📝 Enter command: add Learn Go programming
✅ Added todo: Learn Go programming (ID: 1)

📝 Enter command: add Build a web app
✅ Added todo: Build a web app (ID: 2)

📝 Enter command: list

📋 Your Todo List:
==================
❌ [1] Learn Go programming
❌ [2] Build a web app

📝 Enter command: complete 1
✅ Marked todo 1 as complete: Learn Go programming

📝 Enter command: list

📋 Your Todo List:
==================
✅ [1] Learn Go programming
❌ [2] Build a web app
```

## 📖 Code Structure Explained

### 1. Package Declaration

```go
package main
```

Every Go program starts with a package declaration. `main` is special - it's the entry point.

### 2. Imports

```go
import (
    "bufio"    // For reading user input
    "fmt"      // For printing to screen
    "os"       // For operating system functions
    "strconv"  // For string conversion
    "strings"  // For string manipulation
)
```

### 3. Struct Definition

```go
type Todo struct {
    ID       int    // Unique identifier
    Task     string // The todo description
    Complete bool   // Whether it's done
}
```

Structs group related data together.

### 4. Type Aliases

```go
type TodoList []Todo
```

This creates a new type that's a slice of Todo structs.

### 5. Global Variables

```go
var todos TodoList
var nextID int = 1
```

Variables declared outside functions are accessible throughout the program.

### 6. Functions

Each function demonstrates different Go concepts:

- `addTodo()`: Working with structs and slices
- `listTodos()`: Loops and conditionals
- `completeTodo()`: Modifying data
- `deleteTodo()`: Slice manipulation
- `main()`: Program entry point

## 🔍 Key Go Concepts in This Project

### Variables and Types

- `int`: Whole numbers (1, 2, 3...)
- `string`: Text ("hello", "world")
- `bool`: True/false values
- `[]Todo`: Slice of Todo structs

### Functions

```go
func functionName(parameter type) {
    // function body
}
```

### Control Flow

```go
// If statement
if condition {
    // do something
}

// For loop
for i := 0; i < 10; i++ {
    // loop body
}

// For range (iterate over slices)
for index, item := range slice {
    // loop body
}
```

### Error Handling

```go
value, err := someFunction()
if err != nil {
    // handle error
}
```

## 🎓 Next Steps

After understanding this project, try these exercises:

1. **Add Priority Levels**: Modify the Todo struct to include priority (High, Medium, Low)
2. **Add Due Dates**: Include a due date field for each todo
3. **Save to File**: Make todos persist between program runs
4. **Add Categories**: Group todos by categories (Work, Personal, Shopping)
5. **Search Function**: Add ability to search todos by text

## 📚 Additional Resources

- [Go Tour](https://tour.golang.org/) - Interactive Go tutorial
- [Go by Example](https://gobyexample.com/) - Code examples
- [Effective Go](https://golang.org/doc/effective_go.html) - Go best practices

## 🤝 Contributing

Feel free to modify this project and experiment! That's the best way to learn.

## 📄 License

This project is open source and available under the MIT License.

---

Happy coding! 🎉
