package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	sum := 0
	if numbers == "" {
		return 0, nil
	}

	if strings.Contains(numbers, "//") {
		delimiter := identifyDelimiter(numbers)
		err := calculateSum(numbers, &sum, delimiter)
		if err != nil {
			return 0, err
		}

	} else {
		err := calculateSum(numbers, &sum, ",")
		if err != nil {
			return 0, err
		}
	}
	return sum, nil
}

func identifyDelimiter(numbers string) string {
	values := strings.Split(strings.Trim(numbers, "//"), "\n")
	return values[0]

}

func calculateSum(numbers string, sum *int, delimiter string) error {
	inputs := strings.Split(numbers, delimiter)
	for _, input := range inputs {
		values := strings.Split(input, "\n")
		for _, value := range values {
			n, err := strconv.Atoi(value)
			if n < 0 {
				return identifyNegatives(inputs)
			}
			if err == nil && n < 1000 {
				*sum += n
			} else {
				//TODO
			}
		}
	}
	return nil
}

func identifyNegatives(values []string) error {
	negatives := []int{}
	for _, value := range values {
		n, _ := strconv.Atoi(value)
		if n < 0 {
			negatives = append(negatives, n)
		}
	}
	return errors.New(fmt.Sprint("negatives not allowed: ", negatives))
}
