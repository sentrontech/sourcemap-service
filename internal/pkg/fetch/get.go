package fetch

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const userAgent = `W/"sentron/1.1 (+https://sentron.app/)"`

// Get fetches a URL using the GET method
func Get(url string) ([]byte, error) {
	// create request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", userAgent)

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check status code is 2XX
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.Errorf("Error fetching URL - %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
