package calc

import (
	"errors"
)

type calculator struct {
}

func NewCalculator() *calculator {
	return &calculator{}
}

func (c calculator) Calculate(n1 float64, m string, n2 float64) (float64, error) {
	switch m {
	case "+":
		return add(n1, n2), nil
	case "-":
		return sub(n1, n2), nil
	case "*":
		return mul(n1, n2), nil
	case "/":
		ans, err := div(n1, n2)
		if err != nil {
			return 0, err
		} else {
			return ans, nil
		}
	default:
		return 0, errors.New("несуществующая операция")
	}
}

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func sub(n1, n2 float64) float64 {
	return n1 - n2
}

func mul(n1, n2 float64) float64 {
	return n1 * n2
}

func div(n1, n2 float64) (float64, error) {
	if n2 == 0 {
		return 0, errors.New("на ноль делить нельзя")
	}
	return n1 / n2, nil
}
