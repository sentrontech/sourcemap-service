package code

import (
	"log"

	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/intmath"
)

// MaxLineLength defines the maximum line length allowed
const MaxLineLength = 80

func getStartAndEnd(line string, columnNumber int, maxLineLength int) (start int, end int, truncated bool) {
	lineLength := len(line)

	if columnNumber > lineLength {
		log.Printf("WARNING: Column number greater than line length")
		start = 0
		end = intmath.Min(80, lineLength)
		truncated = (lineLength > 80)
		return
	}

	// Most style guidelines recommend a maximum line length (assumed 80 chars)
	if lineLength <= maxLineLength {
		start = 0
		end = lineLength
		truncated = false
		return
	}

	// Check if the snippet is at the start or the end of the line
	offsetEnd := lineLength - columnNumber
	if offsetEnd > columnNumber {
		start = intmath.Max(columnNumber-40, 0)
		end = intmath.Min(start+80, lineLength)
		truncated = true
		return
	}
	end = intmath.Min(columnNumber+40, lineLength)
	start = intmath.Max(end-80, 0)
	truncated = true
	return
}

// ParseLine takes a line of code, and shortens it if necessary
func ParseLine(line string, columnNumber int) (parsedLine string, truncated bool) {
	// Get start and end position
	start, end, truncated := getStartAndEnd(line, columnNumber, MaxLineLength)
	parsedLine = line[start:end]
	return
}
