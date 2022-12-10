package gofzf

import "testing"

func TestSelect(t *testing.T) {
	selectInput := []string{"swag", "me", "baby"}
	result, err := Select(selectInput)
	if err != nil {
		t.Fatal("err in Select return")
	}
	exists := itemExists(selectInput, result)
	if exists != true {
		t.Fatal("Select result does not exist in selectInput")
	}
}

func itemExists(array []string, item string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return true
		}
	}
	return false
}
