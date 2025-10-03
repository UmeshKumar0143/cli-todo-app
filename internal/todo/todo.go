package todo

import "fmt"


type Item struct {
	Title string
	Done  bool
}

func NewItem(title string) Item {
	return Item{

		Title: title,
		Done:  false,
	}
}


var Todos []Item



func AddNewItem( newItem Item){
		Todos = append(Todos, newItem)
}

func ShowTodos(){
	for i:=0 ; i<len(Todos); i++{
		fmt.Printf("%d. %s done %t  \n" ,i+1,  Todos[i].Title, Todos[i].Done); 
	}
}

func Markdone(index int){
	if index<1 || index>len(Todos){
		fmt.Printf("Enter valid Index %d", index)
	}

	Todos[index-1].Done = true 
	

}

func EditTodo(index int , new_title string )  {
	if index<1 || index>len(Todos){
		fmt.Printf("Enter valid Index %d", index)
	}

	Todos[index-1].Title = new_title
	
	
		
}


