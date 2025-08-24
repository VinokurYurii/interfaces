package pets

type Dog struct {
	Animal
}

func (_d *Dog) Eat(amount uint8) (uint8, error) {
	if amount > 15 {
		return 0, newError("Dog can't eat that much", nil)
	}
	return amount, nil
}

func (_d *Dog) Walk() string {
	return "Dog is walking"
}
