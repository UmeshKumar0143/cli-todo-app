package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Title string
	Done  bool
}

const File_name = "internal/storage/todo.json"

func SaveTodo(Todos []Item) error {
	data, err := json.MarshalIndent(Todos, "", "")

	if err != nil {
		fmt.Println("An Error Occured ")
		return err
	}

	return os.WriteFile(File_name, data, 0644)
}

func LoadTodo() ([]Item, error) {
			data,err := os.ReadFile(File_name)

			if	err!=nil {
				fmt.Println("An error Occured ")
				return nil, err
			}

			var Todos []Item

			if err:=json.Unmarshal(data,&Todos); err!=nil {
				fmt.Println("Error Occured parsing file ")
				return nil,err
			}

			return Todos,nil

}

func NewItem(title string) Item {
	return Item{
		Title: title,
		Done: false ,
	}
}

func AddTodo(Task Item) {
		Todos,err :=  LoadTodo()

		if err!=nil {
			fmt.Println("Error Occured")
			return 
		}

		Todos = append(Todos, Task)
		SaveTodo(Todos)

}

func EditTodo(index int , new_title string) {
		Todos , err :=  LoadTodo()

		if err!=nil {
			fmt.Println("Error Occured")
			return 
		}
		
		if index < 1 || index > len(Todos) {
		fmt.Printf("Enter valid index: %d", index)
		return 
	   }

		Todos[index-1].Title = new_title
		SaveTodo(Todos) 

}

func MarkDone(index int) {

	Todos , err :=  LoadTodo()

    if err!=nil {
	  fmt.Println("Error Occured")
			return 
		}

	if index < 1 || index > len(Todos) {
		fmt.Printf("Enter valid index: %d", index)
		return 
	 }

	 Todos[index-1].Done = true
	 SaveTodo(Todos)

}

func DeleteTodo(index int) {

		Todos , err :=  LoadTodo()

    if err!=nil {
	  fmt.Println("Error Occured")
			return 
		}

	if index < 1 || index > len(Todos) {
		fmt.Printf("Enter valid index: %d", index)
		return 
	 }


	 Todos = append(Todos[:index-1], Todos[index:]... )
	 SaveTodo(Todos)

}

func ShowTodos(){
	
	Todos , err :=  LoadTodo()

    if err!=nil {
	  fmt.Println("Error Occured")
			return 
	}

	if len(Todos) == 0 {
		fmt.Println("No current tasks. Add some tasks...")
		return
	}

	for i, item := range Todos {
		status := " "
		if item.Done {
			status = "✓"
		}
		fmt.Printf("%d. [ %s ] %s\n", i+1, status, item.Title)
	}
}