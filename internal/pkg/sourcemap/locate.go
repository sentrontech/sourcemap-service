package sourcemap

import (
	"fmt"
	"log"
	"strings"

	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/fetch"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/validate"
	"github.com/pkg/errors"
)

func checkSourceMapURLs(jsURL string) (string, error) {
	potentialSourceMapURLs, err := GuessMapURL(jsURL)
	if err != nil {
		return "", err
	}

	for _, sourceMapURL := range potentialSourceMapURLs {
		_, err := fetch.Get(sourceMapURL)
		if err == nil {
			return sourceMapURL, nil
		}
	}

	errMsg := fmt.Sprintf("No sourcemap found at %v", strings.Join(potentialSourceMapURLs, ", "))
	return "", errors.New(errMsg)
}

// Locate returns the source map URL (if one is found) and whether the URL was guessed (based on enumeration)
func Locate(jsFileURL string) (string, bool, error) {
	log.Printf("DEBUG: Validating jsFileURL\n")
	valid, err := validate.URL(jsFileURL)

	if !valid {
		return "", false, err
	}

	log.Printf("DEBUG: Fetching Javascript file\n")
	jsFile, err := fetch.Get(jsFileURL)
	if err != nil {
		return "", false, err
	}

	log.Printf("DEBUG: Parsing source map URL from file\n")
	parsedSourceMapURL, err := ParseURL(jsFileURL, string(jsFile))
	if err != nil {
		return "", false, err
	}

	if parsedSourceMapURL != "" {
		return parsedSourceMapURL, false, nil
	}

	log.Printf("DEBUG: Checking potential sourcemap URLs\n")
	guessedSourceMapURL, err := checkSourceMapURLs(jsFileURL)
	if err != nil {
		return "", false, err
	}

	return guessedSourceMapURL, true, nil
}
