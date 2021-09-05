package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	assert.Equal(t, 81, square(9), "should be 81")
	assert.Equal(t, 9, square(3), "should be 9 ")
}
