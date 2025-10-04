package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/UmeshKumar0143/cli-todo-app.git/internal/todo"
)

func main() {


	args := os.Args[1:]

	reader := bufio.NewReader(os.Stdin)


	switch  args[0] {

		case "add": 

		if len(args) < 2 {
       		 fmt.Println("Error: no task title provided")
       	 	return
    	}


		 title := strings.Join(args[1:], " ")
		 new_todo := todo.NewItem(title)
		 todo.AddTodo(new_todo)
		 fmt.Println("Task Added Successfully")

		case "show": 
			todo.ShowTodos()

		case "edit": 

		 todo_idx,_ := strconv.Atoi(args[1]) 

		 fmt.Println("Enter the New Title")	
		 
		 new_title,_ := reader.ReadString('\n')
		 new_title = strings.TrimSpace(new_title)
		 todo.EditTodo(todo_idx, new_title)

		 todo.ShowTodos()

		case "done": 
		fmt.Println("done command")

		case "delete": 
		fmt.Println("delete command")

	default: 
	fmt.Println("Not a Valid Command ")
	}


	fmt.Println(args)


}