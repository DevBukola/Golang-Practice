package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	ID   int
	Task string
}

var todos []Todo

var fileName = "todosList"

func getExisiting(todos []Todo) []Todo {
	readF, _ := os.ReadFile(fileName)
	var existingTodos []Todo
	json.Unmarshal(readF, &existingTodos)

	return existingTodos
}

func writeTodos(todos []Todo) {
	existingTodos := getExisiting(todos)

	existingTodos = append(existingTodos, todos...)

	data, err := json.MarshalIndent(existingTodos, "", " ")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// defer file.Close()
	os.WriteFile(fileName, data, 0644)

	// for _, todo := range todos {
	// 	fmt.Fprintf(file, "(%v) %s", todo.ID, todo.Task)
	// }
}

func main() {
	reader := bufio.NewReader(os.Stdin)

retry:
	fmt.Print("Enter something: ")
	inp, _ := reader.ReadString('\n')
	args := strings.Split(inp, " ")
	if len(args) < 1 {
		fmt.Println("Please provide an input: 'add','ls', or 'exit'.")
		return
	}

	input := args[0]
	// fmt.Println("INPUT IS",input)
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

		ex := getExisiting([]Todo{})

		id := getLastId(ex) + 1

		todos = []Todo{generateTodo(id, task)}
		writeTodos(todos)
		// fmt.Println(len(todos))
		fmt.Println("Todo:", task)
		goto retry

	case "ls":
		// fmt.Println("i ran")
		fileContent, _ := os.ReadFile(fileName)
		var list []Todo
		json.Unmarshal(fileContent, &list)

		if len(list) == 0 {
			fmt.Println("No todos yet.")
			return
			// continue
		}
		for _, ls := range list {
			fmt.Printf("%v. %s\n", ls.ID, ls.Task)

		}
	case "exit":
		{
			fmt.Println("Goodbye.")
			return
		}

	default:
		fmt.Println("Please use add or ls.")
	}

	// writeTodos(todos)
}

func generateTodo(id int, task string) Todo {
	return Todo{
		ID:   id,
		Task: task,
	}

}

func getLastId(el []Todo) int {
	fmt.Println("el:",el)
	if len(el) < 1 {
		return 0
	}

	return el[len(el)-1].ID
}
