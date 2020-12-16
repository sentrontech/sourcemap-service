package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_URL_Empty(t *testing.T) {
	valid, err := URL("")
	assert.Equal(t, false, valid)
	assert.EqualErrorf(t, err, "URL is not defined", "Error string mismatch")
}

func Test_URL_InvalidScheme(t *testing.T) {
	valid, err := URL("mongodb://example.com")
	assert.Equal(t, false, valid)
	assert.EqualErrorf(t, err, "URL scheme must be 'http' or 'https'", "Error string mismatch")
}

func Test_URL_RelativeUrl(t *testing.T) {
	valid, err := URL("./foo/bar")
	assert.Equal(t, false, valid)
	assert.EqualErrorf(t, err, "Invalid URL", "Error string mismatch")
}

func Test_URL_AbsoluteUrl(t *testing.T) {
	valid, err := URL("https://google.com")
	assert.Equal(t, true, valid)
	assert.Nil(t, err)
}
func Test_URL_IpAddress(t *testing.T) {
	valid, err := URL("https://1.2.3.4")
	assert.Equal(t, true, valid)
	assert.Nil(t, err)
}

func Test_URL_HostWithPort(t *testing.T) {
	valid, err := URL("https://host:3000")
	assert.Equal(t, true, valid)
	assert.Nil(t, err)
}
