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

		case "list": 
			todo.ShowTodos()

		case "edit": 

		 todo_idx,_ := strconv.Atoi(args[1]) 

		 fmt.Print("Enter the New Title: ")	
		 
		 new_title,_ := reader.ReadString('\n')
		 new_title = strings.TrimSpace(new_title)
		 todo.EditTodo(todo_idx, new_title)

		fmt.Printf(" Task %d  Edited Succesfuly \n", todo_idx)


		 todo.ShowTodos()

		case "done": 

			todo_idx,_ := strconv.Atoi(args[1])

			todo.MarkDone(todo_idx)

			fmt.Printf("%d Task Marked as Complete \n", todo_idx)

			todo.ShowTodos()

		case "delete": 

			todo_idx,_ := strconv.Atoi(args[1])

			todo.DeleteTodo(todo_idx)

			fmt.Println("Task Deleted Succesfully ")

			todo.ShowTodos()
	case "deleteAll":
		todo.DeleteAll()
		fmt.Println("All Todos Deleted Succesfully ")
	case "-h","--help","help":
		todo.PrintHelp()

	default: 
	fmt.Println("Not a Valid Command ")
	}




}