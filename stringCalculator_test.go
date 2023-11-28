package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddWithInputEmptyString(t *testing.T) {
	sum, _ := Add("")
	assert.Equal(t, 0, sum)
}

func TestAddWithInputOneNumber(t *testing.T) {
	sum, _ := Add("4")
	assert.Equal(t, 4, sum)
}

func TestAddWithIputTwoNumbers(t *testing.T) {
	sum, _ := Add("9,7")
	assert.Equal(t, 16, sum)
}

func TestAddWithInputThreeNumbers(t *testing.T) {
	sum, _ := Add("2,3,2")
	assert.Equal(t, 7, sum)
}

func TestAddWithNewLineInInput(t *testing.T) {
	sum, _ := Add("1\n2,3")
	assert.Equal(t, 6, sum)
}

func TestAddWithDifferentDelimiters(t *testing.T) {
	sum, _ := Add("//;\n1;2")
	assert.Equal(t, 3, sum)

}

func TestIdentifyDelimiter(t *testing.T) {
	assert.Equal(t, ";", identifyDelimiter("//;\n1;2"))
}

func TestAddWithDifferentDelimiters2(t *testing.T) {
	sum, _ := Add("//***\n1***2***3")
	assert.Equal(t, 6, sum)

}

func TestAddWithNegativeNumberInInput(t *testing.T) {
	_, err := Add("2,-2,-4")

	expectedErr := "negatives not allowed: [-2 -4]"
	assert.EqualError(t, err, expectedErr)
}

func TestAddWithInputBiggerThanThousand(t *testing.T) {
	sum, _ := Add("2,1001")
	assert.Equal(t, 2, sum)
}
