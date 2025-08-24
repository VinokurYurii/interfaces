package pets

type Dog struct {
	Name string
}

func (c *Dog) Eat(amount uint8) (uint8, error) {
	if amount > 15 {
		return 0, newError("Dog can't eat that much", nil)
	}
	return amount, nil
}
