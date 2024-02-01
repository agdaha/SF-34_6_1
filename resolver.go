package main

import (
	"errors"
	"fmt"
	"regexp"
	"skillfactory/SF-34_6_1/calc"
	"strconv"
)

func Resolver(l []string) ([]string, error) {
	if len(l) == 0 {
		return nil, errors.New("на входе функции пустой список")
	}
	result := []string{}

	re := regexp.MustCompile(`(\d+)([+-\/\*])(\d+)[=]([\?])`)
	calc := calc.NewCalculator()

	for _, s := range l {
		groups := re.FindStringSubmatch(s)
		if len(groups) != 5 {
			continue
		}
		n1, err := strconv.ParseFloat(groups[1], 64)
		if err != nil {
			continue
		}
		op := groups[2]
		n2, err := strconv.ParseFloat(groups[3], 64)
		if err != nil {
			continue
		}
		res, err := calc.Calculate(n1, op, n2)
		if err != nil {
			continue
		}
		result = append(result, fmt.Sprintf("%v%v%v=%v", n1, op, n2, res))
	}

	return result, nil
}
