package sourcemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func Test_Locate_FileNotFound(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/dist/main.js").
		Reply(404).
		JSON(map[string]string{"message": "Not found"})

	sourceMapURL, isGuess, err := Locate("http://example.com/dist/main.js")
	assert.Equal(t, "", sourceMapURL)
	assert.Equal(t, false, isGuess)
	assert.EqualErrorf(t, err, "Error fetching URL - 404", "Error message mismatch")
}

func Test_Locate_AbsUrl(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/dist/main.js").
		Reply(200).
		File("../../../test/data/abs-path.js")

	sourceMapURL, isGuess, err := Locate("http://example.com/dist/main.js")
	assert.Equal(t, "https://cdn.example.com/main.js.map", sourceMapURL)
	assert.Equal(t, false, isGuess)
	assert.Nil(t, err)
}
func Test_Locate_RelativeUrl(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/dist/main.js").
		Reply(200).
		File("../../../test/data/relative-path.js")

	sourceMapURL, isGuess, err := Locate("http://example.com/dist/main.js")
	assert.Equal(t, "http://example.com/sourcemaps/main.js.map", sourceMapURL)
	assert.Equal(t, false, isGuess)
	assert.Nil(t, err)
}

func Test_Locate_GuessWithReplacement(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/public/app.min.js").
		Reply(200).
		BodyString("")

	gock.New("http://example.com").
		Get("/public/app.min.map").
		Reply(200).
		BodyString("")

	gock.New("http://example.com").
		Get("/public/app.min.js.map").
		Reply(404).
		BodyString("")

	sourceMapURL, isGuess, err := Locate("http://example.com/public/app.min.js")
	assert.Equal(t, "http://example.com/public/app.min.map", sourceMapURL)
	assert.Equal(t, true, isGuess)
	assert.Nil(t, err)
}
func Test_Locate_GuessWithSuffix(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/dist/main.js").
		Reply(200).
		BodyString("")

	gock.New("http://example.com").
		Get("/dist/main.map").
		Reply(404).
		BodyString("")

	gock.New("http://example.com").
		Get("/dist/main.js.map").
		Reply(200).
		BodyString("")

	sourceMapURL, isGuess, err := Locate("http://example.com/dist/main.js")
	assert.Equal(t, "http://example.com/dist/main.js.map", sourceMapURL)
	assert.Equal(t, true, isGuess)
	assert.Nil(t, err)
}
