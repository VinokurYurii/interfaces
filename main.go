package main

import (
	"fmt"
	"interfaces/pets"
	"os"
)

func main() {
	cat := pets.Cat{Name: "mr. buttons"}
	dog := pets.Dog{Name: "spot"}

	var feedToCat uint8 = 13
	var feedToDog uint8 = 8

	catFed, err := feed(&cat, feedToCat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding cat: : %v\n", err)
	} else {
		fmt.Println("Cat ate:", catFed)
	}

	fmt.Print("\n\n\t =====\n\n\n")

	dogFed, err := feed(&dog, feedToDog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding dog: : %v\n", err)
	} else {
		fmt.Println("Dog ate:", dogFed)
	}
}

func feed(animal pets.Eater, amount uint8) (uint8, error) {
	return animal.Eat(amount)
}
