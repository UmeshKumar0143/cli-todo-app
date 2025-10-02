package todo

import "testing"

func NewItemTest(t *testing.T) {
	item := NewItem("Hello")

	if item.title != "Hello" {
		t.Errorf("Expected string Hello , recived %s", item.title)
	}

}
