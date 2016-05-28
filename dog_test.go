package main

import (
	"testing"
)

func TestDogSay(t *testing.T) {
	dog := Dog{"Marley", 10, "Woof woof!"}

	if dog.say() != "Woof woof!" {
		t.Error("Expected Woof woof!, got", dog.say())
	}
}
