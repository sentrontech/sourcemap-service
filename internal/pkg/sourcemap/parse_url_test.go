package sourcemap

import (
	"testing"

	"github.com/jpstevens/sentron-sourcemaps/test/helper"
	"github.com/stretchr/testify/assert"
)

func Test_ParseURL_RelativePathSameDir(t *testing.T) {
	jsFile := string(helper.LoadTestData("coffee-min.js"))
	sourceURL, err := ParseURL("http://example.com/assets/js/main.js", jsFile)
	assert.Equal(t, "http://example.com/assets/js/main.js.map", sourceURL)
	assert.Nil(t, err)
}

func Test_ParseURL_RelativePathDifferentDir(t *testing.T) {
	jsFile := string(helper.LoadTestData("relative-path.js"))
	sourceURL, err := ParseURL("http://example.com/assets/js/main.js", jsFile)
	assert.Equal(t, "http://example.com/assets/sourcemaps/main.js.map", sourceURL)
	assert.Nil(t, err)
}

func Test_ParseURL_AbsolutePath(t *testing.T) {
	jsFile := string(helper.LoadTestData("abs-path.js"))
	sourceURL, err := ParseURL("http://example.com/assets/js/main.js", jsFile)
	assert.Equal(t, "https://cdn.example.com/main.js.map", sourceURL)
	assert.Nil(t, err)
}

func Test_ParseURL_AbsolutePathNoProtocol(t *testing.T) {
	jsFile := string(helper.LoadTestData("abs-path-no-protocol.js"))
	sourceURL, err := ParseURL("http://example.com/assets/js/main.js", jsFile)
	assert.Equal(t, "http://cdn.example.com/main.js.map", sourceURL)
	assert.Nil(t, err)
}

func Test_ParseURL_NotFound(t *testing.T) {
	jsFile := string(helper.LoadTestData("small-file.js"))
	sourceURL, err := ParseURL("http://example.com/assets/js/main.js", jsFile)
	assert.Equal(t, "", sourceURL)
	assert.Nil(t, err)
}

func benchmarkParseURL(b *testing.B, filePath string) {
	jsFile := string(helper.LoadTestData(filePath))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseURL("http://example.com", jsFile)
	}
}

func Benchmark_ParseURL_Jquery(b *testing.B)    { benchmarkParseURL(b, "jquery.min.js") }
func Benchmark_ParseURL_Bootstrap(b *testing.B) { benchmarkParseURL(b, "bootstrap.min.js") }
func Benchmark_ParseURL_Datadog(b *testing.B)   { benchmarkParseURL(b, "dd.js") }
