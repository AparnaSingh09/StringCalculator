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
		delimiters := identifyDelimiters(numbers)
		err := calculateSum(numbers, &sum, delimiters)
		if err != nil {
			return 0, err
		}

	} else {
		err := calculateSum(numbers, &sum, []string{","})
		if err != nil {
			return 0, err
		}
	}
	return sum, nil
}

func identifyDelimiters(numbers string) []string {
	delimiters := []string{}
	values := strings.Split(strings.Trim(numbers, "//"), "\n")
	va := strings.Split(values[0], "[")
	for _, j := range va {
		if j != "" {
			delimiters = append(delimiters, strings.Split(j, "]")[0])
		}

	}
	fmt.Println(delimiters)
	return delimiters

}

func splitAny(s string, delimiters []string) []string {
	seps := strings.Join(delimiters, "")
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func calculateSum(numbers string, sum *int, delimiters []string) error {

	inputs := splitAny(numbers, delimiters)
	inputArr := []string{}
	for _, input := range inputs {
		strArr := strings.Split(input, "\n")
		for _, str := range strArr {
			inputArr = append(inputArr, str)
		}
	}

	for _, value := range inputArr {
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
