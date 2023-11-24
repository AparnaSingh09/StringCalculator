package main

import "strconv"

func Add(numbers string) int {
	var sum int
	switch len(numbers) {
	case 0:
		sum = 0
	case 1:
		n, err := strconv.Atoi(numbers)
		if err == nil {
			sum = n
		} else {
			//TODO
		}

	}
	return sum
}
