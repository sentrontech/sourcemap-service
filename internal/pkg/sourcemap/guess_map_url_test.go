package sourcemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GuessMapURL_NotHttpsOrHttp(t *testing.T) {
	urls, err := GuessMapURL("mongodb://localhost")
	expErr := "URL scheme must be 'http' or 'https'"
	assert.Nil(t, urls)
	assert.EqualErrorf(t, err, expErr, "Error message mismatch")
}

func Test_GuessMapURL_NotUrl(t *testing.T) {
	urls, err := GuessMapURL("some-file.js")
	expErr := "Cannot parse relative URL"
	assert.Nil(t, urls)
	assert.EqualErrorf(t, err, expErr, "Error message mismatch")
}

func Test_GuessMapURL_NotJsExt(t *testing.T) {
	urls, err := GuessMapURL("https://example.com/test.coffee")
	expErr := "URL path must end with '.js' extension"
	assert.Nil(t, urls)
	assert.EqualErrorf(t, err, expErr, "Error message mismatch")
}

func Test_GuessMapURL_MissingExt(t *testing.T) {
	urls, err := GuessMapURL("https://example.com/js")
	expErr := "URL path must end with '.js' extension"
	assert.Nil(t, urls)
	assert.EqualErrorf(t, err, expErr, "Error message mismatch")
}

func Test_GuessMapURL_BasicUrl(t *testing.T) {
	urls, err := GuessMapURL("https://example.com/test.js")
	expUrls := []string{"https://example.com/test.js.map", "https://example.com/test.map"}
	assert.Equal(t, expUrls, urls)
	assert.Nil(t, err)
}

func Test_GuessMapURL_MinJsExt(t *testing.T) {
	urls, err := GuessMapURL("https://example.com/test.min.js")
	expUrls := []string{"https://example.com/test.min.js.map", "https://example.com/test.min.map"}
	assert.Equal(t, expUrls, urls)
	assert.Nil(t, err)
}

func Test_GuessMapURL_UrlWithQuery(t *testing.T) {
	urls, err := GuessMapURL("https://example.com/test.js?example=1")
	expUrls := []string{"https://example.com/test.js.map", "https://example.com/test.map"}
	assert.Equal(t, expUrls, urls)
	assert.Nil(t, err)
}

func Test_GuessMapURL_UrlWithFragment(t *testing.T) {
	urls, err := GuessMapURL("https://example.com/test.js#fragment")
	expUrls := []string{"https://example.com/test.js.map", "https://example.com/test.map"}
	assert.Equal(t, expUrls, urls)
	assert.Nil(t, err)
}
