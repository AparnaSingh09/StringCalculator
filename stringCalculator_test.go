package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddWithInputEmptyString(t *testing.T) {
	assert.Equal(t, 0, Add(""))
}

func TestAddWithInputOneNumber(t *testing.T) {
	assert.Equal(t, 4, Add("4"))
}

func TestAddWithIputTwoNumbers(t *testing.T) {
	assert.Equal(t, 16, Add("9,7"))
}

func TestAddWithInputThreeNumbers(t *testing.T) {
	assert.Equal(t, 7, Add("2,3,2"))
}
