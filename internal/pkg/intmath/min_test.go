package intmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Min_FirstSmaller(t *testing.T) {
	received := Min(99, 1)
	expected := 1
	assert.Equal(t, expected, received)
}

func Test_Min_SecondSmaller(t *testing.T) {
	received := Min(99, 101)
	expected := 99
	assert.Equal(t, expected, received)
}

func Test_Min_BothEqual(t *testing.T) {
	received := Min(3000, 3000)
	expected := 3000
	assert.Equal(t, expected, received)
}
