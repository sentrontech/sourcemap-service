package sourcemap

import (
	"fmt"

	"github.com/go-sourcemap/sourcemap"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/code"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/fetch"
)

// GenerateExtract returns a snippet of code re-formed from the sourcemap data
func GenerateExtract(mapURL string, line int, column int) (
	_ string,
	_ string,
	_ int,
	_ int,
	_ []code.LineExtract,
	err error,
) {
	// fetch sourcemap file
	smapFile, err := fetch.Get(mapURL)
	if err != nil {
		err = fmt.Errorf("Cannot fetch source map: %s", err)
		return
	}

	// parse sourcemap
	smapConsumer, err := sourcemap.Parse(mapURL, smapFile)
	if err != nil {
		err = fmt.Errorf("Cannot parse source map: %s", err)
		return
	}

	// find the source
	fileURL, functionName, lineNumber, columnNumber, ok := smapConsumer.Source(line, column)
	if !ok {
		err = fmt.Errorf("Cannot consume source map")
		return
	}

	// fetch the src file
	srcFile, err := fetch.Get(fileURL)
	if err != nil {
		err = fmt.Errorf("Cannot fetch source file: %s", err)
		return
	}

	// parse src file extract
	extract := code.Extract(srcFile, lineNumber, columnNumber)

	// create response
	return fileURL, functionName, lineNumber, columnNumber, extract, nil
}
