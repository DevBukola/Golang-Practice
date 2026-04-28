package main

import (
	// "bufio"

	"bufio"
	// "encoding/json"
	"fmt"
	"os"
	"strings"
	// "strings"
)

type Todo struct {
	ID int `json:""`
	Task string `json:""`
}

var todos []Todo

var fileName = "todosList"

func writeTodos(todos []Todo) {
	// data, err := json.MarshalIndent(todos, "", " ")
	file, err:= os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// os.WriteFile(fileName, file, 0644)
	// defer file.Close()

	for _, todo := range todos {
		fmt.Fprintf(file, "(%v) %s", todo.ID, todo.Task)
	}
}



func main() {
	reader := bufio.NewReader(os.Stdin)

	retry:
	fmt.Print("Enter something: ")
	inp, _ := reader.ReadString('\n')
	args := strings.Split(inp," ")
	if len(args) < 1 {
		fmt.Println("Please provide an input: 'add','list', or 'exit'.")
		return
	}

	input := args[0]
	// fmt.Println("n : ",input)
	switch strings.TrimSpace(input) {
	case "add":
		if len(args) < 2 {
			fmt.Println("Kindly, enter a todo task.")
			return
		}

		// task := args[1]
		task := strings.Join(args[1:], " ")
		// task = strings.TrimSpace(task)
		// id := len(todos+1)
		id := len(todos) + 1

		todos = append(todos, generateTodo(id, task))
		writeTodos(todos)
		// fmt.Println(len(todos))
		fmt.Println("Todo:", task)
		goto retry

	case "list":
		// fmt.Println("i ran")
		if len(todos) == 0 {
			fmt.Println("No todos yet.")
			return
			// continue
		}
		for _, todo := range todos {
			fmt.Printf("%v. %s\n", todo.ID, todo.Task)
		// writeTodos()
		List, _ := os.ReadFile(fileName)
		fmt.Printf(string(List))
		return
	
	}
	case "exit": {
	fmt.Println("Goodbye.")
	return
	}

	default:
		fmt.Println("Please use add or list.")
	}

	writeTodos(todos)	
}



func generateTodo(id int, task string) Todo {
	return Todo{
		ID:   id,
		Task: task,
	}
	
}


// func main() {
// 	var todos []Todo

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter todo:")
	
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error", err)
// 		return
// 	}

// 	// newTodo := Todo {
// 	// 	ID: 1,
// 	// 	Task: input,
// 	// }
// 	todos = append(todos, generateTodo(1, input))

// 	fmt.Println("Todo:",todos)
// }

// func generateTodo(id int, task string) Todo{
// 	return Todo {
// 		ID: id,
// 		Task: task,
// 	}
// }
