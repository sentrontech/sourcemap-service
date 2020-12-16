package code

import (
	"strings"

	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/intmath"
)

// NumLinesBefore defines the max number of lines to show before the provided line number
const NumLinesBefore = 5

// NumLinesAfter defines the max number of lines to show after the provided line number
const NumLinesAfter = 5

// LineExtract contains the line number, line content and whether the line has been shorted (i.e. is a snippet)
type LineExtract struct {
	Number    int    `json:"number"`
	Content   string `json:"content"`
	Truncated bool   `json:"truncated"`
}

func makeLineExtract(i int, line string, focusLineNumber int, focusColumnNumber int) LineExtract {
	currentLineNumber := i + 1
	currentColumnNumber := 0
	if currentLineNumber == focusLineNumber {
		currentColumnNumber = focusColumnNumber
	}
	content, truncated := ParseLine(line, currentColumnNumber)
	return LineExtract{
		Number:    currentLineNumber,
		Content:   content,
		Truncated: truncated,
	}
}

// Extract can be run using either:
// - the source JS file, or
// - the minified file, if no sourcemaps can be located
func Extract(srcFile []byte, lineNumber int, columnNumber int) []LineExtract {
	lines := strings.Split(string(srcFile), "\n")
	start := intmath.Max(lineNumber-NumLinesBefore, 1)
	end := intmath.Min(lineNumber+NumLinesAfter, len(lines))
	output := []LineExtract{}
	for i := start - 1; i < end; i++ {
		lineExtract := makeLineExtract(i, lines[i], lineNumber, columnNumber)
		output = append(output, lineExtract)
	}
	return output
}
