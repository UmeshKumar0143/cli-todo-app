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
		

}

func EditTodo() {

}

func MarkDone() {

}

func DeleteTodo() {

}
