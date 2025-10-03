package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/UmeshKumar0143/cli-todo-app.git/internal/todo"
)

func main() {

	var n int 
	fmt.Println("Enter the number of todos to add ")
	fmt.Scanf("%d", &n)

	for i:=0;  i<n; i++ {
		reader := bufio.NewReader(os.Stdin)
		todo_tittle,_ := reader.ReadString('\n')
		todo_tittle = strings.TrimSpace(todo_tittle)
		item := todo.NewItem(todo_tittle) 	
		todo.AddNewItem(item)
	}

	
	todo.ShowTodos()

	fmt.Println("Enter task to be marked Completed ")
     var idx int 
	 fmt.Scanf("%d", &idx)

	 todo.Markdone(idx)
	 
	 todo.ShowTodos()






}