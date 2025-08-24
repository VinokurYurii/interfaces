package main

import (
	"fmt"
	"interfaces/pets"
	"os"
)

func main() {
	cat := pets.Cat{pets.Animal{Name: "mr. buttons"}}
	dog := pets.Dog{pets.Animal{Name: "spot"}}

	var feedToCat uint8 = 13
	var feedToDog uint8 = 8

	catFed, err := feed(&cat, feedToCat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding cat: : %v\n", err)
	} else {
		fmt.Println("Cat ate:", catFed)
	}

	fmt.Printf("Obj structure: %+v", cat)
	fmt.Print("\n\n\t =====\n\n\n")

	dogFed, err := feed(&dog, feedToDog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding dog: : %v\n", err)
	} else {
		fmt.Println("Dog ate:", dogFed)
	}
	fmt.Printf("Obj structure: %+v", dog)
}

func feed(animal pets.EaterWalkerNamed, amount uint8) (uint8, error) {
	fmt.Println("First, let's walk!")
	fmt.Println(animal.Walk())
	fmt.Println("Now, let's feed our", animal.GetName())
	return animal.Eat(amount)
}
