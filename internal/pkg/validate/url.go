package validate

import (
	"net/url"

	"github.com/pkg/errors"
)

// URL checks if a given string is a valid http/https URL
func URL(str string) (bool, error) {
	if str == "" {
		return false, errors.New("URL is not defined")
	}
	u, err := url.ParseRequestURI(str)
	if err != nil {
		return false, errors.New("Invalid URL")
	}
	if u.Scheme == "" {
		return false, errors.New("Missing URL scheme")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false, errors.New("URL scheme must be 'http' or 'https'")
	}
	if u.Host == "" {
		return false, errors.New("Missing host")
	}
	return true, nil
}
