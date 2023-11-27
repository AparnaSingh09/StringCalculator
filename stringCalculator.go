package main

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	sum := 0
	if numbers == "" {
		return 0
	}
	inputs := strings.Split(numbers, ",")
	for _, input := range inputs {
		if strings.Contains(input, "\n") {
			values := strings.Split(input, "\n")
			for _, value := range values {
				calculateSum(value, &sum)
			}
		} else {
			calculateSum(input, &sum)
		}
	}
	return sum
}

func calculateSum(value string, sum *int) {
	n, err := strconv.Atoi(value)
	if err == nil {
		*sum += n
	} else {
		//TODO
	}
}
