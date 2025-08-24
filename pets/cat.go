package pets

type Cat struct {
	Animal
}

func (_c *Cat) Eat(amount uint8) (uint8, error) {
	if amount > 5 {
		return 0, newError("Cat can't eat that much", nil)
	}
	return amount, nil
}

func (_c *Cat) Walk() string {
	return "Cat is walking"
}
