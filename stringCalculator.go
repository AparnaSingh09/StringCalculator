package main

import (
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	sum := 0
	if numbers == "" {
		return 0, nil
	}

	if strings.Contains(numbers, "//") {
		delimiter := IdentifyDelimiter(numbers)
		calculateSum(numbers, &sum, delimiter)

	} else {
		calculateSum(numbers, &sum, ",")
	}
	return sum, nil
}

func IdentifyDelimiter(numbers string) string {
	values := strings.Split(strings.Trim(numbers, "//"), "\n")
	return values[0]

}

func calculateSum(numbers string, sum *int, delimiter string) {
	inputs := strings.Split(numbers, delimiter)
	for _, input := range inputs {
		values := strings.Split(input, "\n")
		for _, value := range values {
			n, err := strconv.Atoi(value)
			if err == nil {
				*sum += n
			} else {
				//TODO
			}
		}
	}
}
