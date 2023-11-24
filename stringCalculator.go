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
		n, err := strconv.Atoi(input)
		if err == nil {
			sum += n
		} else {
			//TODO
		}

	}
	return sum
}
