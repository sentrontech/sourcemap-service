package code

import (
	"testing"

	"github.com/jpstevens/sentron-sourcemaps/test/helper"
	"github.com/stretchr/testify/assert"
)

func Test_Extract_SmallFile(t *testing.T) {
	lineNumber, columnNumber := 3, 3
	testData := helper.LoadTestData("small-file.js")

	received := Extract(testData, lineNumber, columnNumber)
	expected := []LineExtract{
		{Number: 1, Content: "// this is my best code ever!", Truncated: false},
		{Number: 2, Content: "function test() {", Truncated: false},
		{Number: 3, Content: "  throw new Error(\"Oops...\");", Truncated: false},
		{Number: 4, Content: "}", Truncated: false},
		{Number: 5, Content: "if (window.test) {", Truncated: false},
		{Number: 6, Content: "  test();", Truncated: false},
		{Number: 7, Content: "}", Truncated: false},
	}
	assert.Equal(t, expected, received)
}
func Test_Extract_BigFile(t *testing.T) {
	lineNumber, columnNumber := 23, 44
	testData := helper.LoadTestData("big-file.js")

	received := Extract(testData, lineNumber, columnNumber)
	expected := []LineExtract{
		{Number: 18, Truncated: false, Content: "        this.value += increment"},
		{Number: 19, Truncated: false, Content: "        return this"},
		{Number: 20, Truncated: false, Content: "    }"},
		{Number: 21, Truncated: false, Content: "    set value (val) {"},
		{Number: 22, Truncated: true, Content: "        // this is a very long comment so we only expect the first 80 characters"},
		{Number: 23, Truncated: false, Content: "        if (typeof val !== \"number\") throw new Error(\"Value must be a number\")"},
		{Number: 24, Truncated: false, Content: "        this.__value = val"},
		{Number: 25, Truncated: false, Content: "    }"},
		{Number: 26, Truncated: false, Content: "    get value () {"},
		{Number: 27, Truncated: false, Content: "        return this.__value"},
		{Number: 28, Truncated: false, Content: "    }"},
	}
	assert.Equal(t, expected, received)
}

func Test_Extract_SingleLine(t *testing.T) {
	lineNumber, columnNumber := 1, 1
	testData := []byte("test")

	received := Extract(testData, lineNumber, columnNumber)
	expected := []LineExtract{
		{Number: 1, Truncated: false, Content: "test"},
	}
	assert.Equal(t, expected, received)
}

func Test_Extract_OutOfBounds(t *testing.T) {
	lineNumber, columnNumber := 999, 999
	testData := []byte{}

	received := Extract(testData, lineNumber, columnNumber)
	expected := []LineExtract{}
	assert.Equal(t, expected, received)
}
