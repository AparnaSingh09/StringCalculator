package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddWithInputEmptyString(t *testing.T) {
	assert.Equal(t, 0, Add(""))
}
