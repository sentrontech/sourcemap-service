package intmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Max_FirstBigger(t *testing.T) {
	received := Max(99, 1)
	expected := 99
	assert.Equal(t, expected, received)
}

func Test_Max_SecondBigger(t *testing.T) {
	received := Max(99, 101)
	expected := 101
	assert.Equal(t, expected, received)
}

func Test_Max_BothEqual(t *testing.T) {
	received := Max(3000, 3000)
	expected := 3000
	assert.Equal(t, expected, received)
}
