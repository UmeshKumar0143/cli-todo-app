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

func LoadTodo() {

}

func NewItem(title string) Item {
	return Item{}
}

func AddTodo(Task Item) {

}

func EditTodo() {

}

func MarkDone() {

}

func DeleteTodo() {

}
