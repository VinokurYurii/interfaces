package pets

import "fmt"

type Eater interface {
	Eat(amount uint8) (uint8, error)
}

type Walker interface {
	Walk() string
}

type EaterWalker interface {
	Eater
	Walker
}

func newError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}

	return fmt.Errorf(msg)
}
