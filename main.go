package main

import (
	"fmt"
	"interfaces/pets"
	"os"
)

func main() {
	cat := pets.Cat{
		Animal:   pets.Animal{Name: "mr. buttons"},
		Age:      5,
		IsAsleep: true,
	}
	dog := pets.Dog{
		Animal:  pets.Animal{Name: "spot"},
		WWeight: 30,
		Age:     6,
	}

	var feedToCat uint8 = 33
	var feedToDog uint8 = 8

	catFed, err := feed(&cat, feedToCat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding cat: %v\n", err)
	} else {
		fmt.Println("Cat ate:", catFed)
	}

	fmt.Printf("Obj structure: %+v", cat)
	fmt.Print("\n\n\t =====\n\n\n")

	dogFed, err := feed(&dog, feedToDog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding dog: %v\n", err)
	} else {
		fmt.Println("Dog ate:", dogFed)
	}
	fmt.Printf("Obj structure: %+v", dog)
	fmt.Print("\n\n\t =====\n\n\n")
	displayInfo(cat)
	displayInfo(dog)
	displayInfo(42)
	displayInfo(map[string]int{"ololo": 15})
}

func feed(animal pets.EaterWalkerNamed, amount uint8) (uint8, error) {
	switch v := animal.(type) {
	case *pets.Cat:
		fmt.Println(v.GetName(), "its a cat aged", v.Age)
	case *pets.Dog:
		fmt.Println(v.GetName(), "its a dog aged", v.Age)
	default:
		fmt.Println("Unknown animal type")
	}
	fmt.Println("First, let's walk!")
	fmt.Println(animal.Walk())
	fmt.Println("Now, let's feed our", animal.GetName())
	return animal.Eat(amount)
}

func displayInfo(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("This is a string:", v)
	case int:
		fmt.Println("This is a int:", v)
	case pets.Cat:
		fmt.Println("This is a cat named:", v.GetName(), "and it is", v.Age, "years old")
	case pets.Dog:
		fmt.Println("This is a dog named:", v.GetName(), "and weight:", v.WWeight)
	default:
		fmt.Println("Unknow type")
	}
}
