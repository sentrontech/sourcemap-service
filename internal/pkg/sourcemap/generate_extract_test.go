package sourcemap

import (
	"testing"

	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/code"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func Test_GenerateExtract_SourceMapNotFound(t *testing.T) {
	defer gock.Off()

	mapURL := "https://example.com/js/main.js"

	gock.New(mapURL).
		Reply(404).
		BodyString("Not found")

	_, _, _, _, _, err := GenerateExtract(mapURL, 1, 1)
	assert.EqualErrorf(t, err, "Cannot fetch source map: Error fetching URL - 404", "Error message mismatch")
}

func Test_GenerateExtract_InvalidSourceMap(t *testing.T) {
	defer gock.Off()

	mapURL := "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.map"

	gock.New(mapURL).
		Reply(200).
		BodyString("I am invalid")

	_, _, _, _, _, err := GenerateExtract(mapURL, 2, 147)
	assert.EqualErrorf(t, err, "Cannot parse source map: invalid character 'I' looking for beginning of value", "Error message mismatch")
}

func Test_GenerateExtract_CannotConsume(t *testing.T) {
	defer gock.Off()

	mapURL := "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.map"

	gock.New(mapURL).
		Reply(200).
		File("../../../test/data/jquery.min.map")

	_, _, _, _, _, err := GenerateExtract(mapURL, 9999, 9999)
	assert.EqualErrorf(t, err, "Cannot consume source map", "Error message mismatch")
}

func Test_GenerateExtract_CannotFetchSourceFile(t *testing.T) {
	defer gock.Off()

	mapURL := "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.map"
	jsFileURL := "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.js"

	gock.New(mapURL).
		Reply(200).
		File("../../../test/data/jquery.min.map")

	gock.New(jsFileURL).
		Reply(404).
		BodyString("Not found")

	_, _, _, _, _, err := GenerateExtract(mapURL, 2, 27878)
	assert.EqualErrorf(t, err, "Cannot fetch source file: Error fetching URL - 404", "Error message mismatch")
}

func Test_GenerateExtract_Jquery(t *testing.T) {
	defer gock.Off()

	mapURL := "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.map"
	expFileURL := "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.js"
	expLine := 3652
	expCol := 1
	expLineExtract := []code.LineExtract{
		{Number: 3647, Truncated: false, Content: ""},
		{Number: 3648, Truncated: false, Content: "function Identity( v ) {"},
		{Number: 3649, Truncated: false, Content: "\treturn v;"},
		{Number: 3650, Truncated: false, Content: "}"},
		{Number: 3651, Truncated: false, Content: "function Thrower( ex ) {"},
		{Number: 3652, Truncated: false, Content: "\tthrow ex;"},
		{Number: 3653, Truncated: false, Content: "}"},
		{Number: 3654, Truncated: false, Content: ""},
		{Number: 3655, Truncated: false, Content: "function adoptValue( value, resolve, reject, noValue ) {"},
		{Number: 3656, Truncated: false, Content: "\tvar method;"},
		{Number: 3657, Truncated: false, Content: ""},
	}

	gock.New(mapURL).
		Reply(200).
		File("../../../test/data/jquery.min.map")

	gock.New(expFileURL).
		Reply(200).
		File("../../../test/data/jquery.js")

	fileURL, fn, line, col, extract, err := GenerateExtract(mapURL, 2, 27878)

	assert.Equal(t, expFileURL, fileURL)
	assert.Equal(t, "", fn)
	assert.Equal(t, expLine, line)
	assert.Equal(t, expCol, col)
	assert.Equal(t, expLineExtract, extract)
	assert.Nil(t, err)
}
