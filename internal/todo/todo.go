package todo 

type Item struct {
	title string 
	done bool 
}


func NewItem(title string) Item {
	return Item{
		title: title,
		done: false,
	}
}