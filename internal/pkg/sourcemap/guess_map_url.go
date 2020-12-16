package sourcemap

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

const suffix = ".js"

// GuessMapURL guesses sourcemap URLs
func GuessMapURL(jsURL string) (urls []string, err error) {
	u, err := url.Parse(jsURL)
	if err != nil {
		return
	}
	if !u.IsAbs() {
		err = errors.New("Cannot parse relative URL")
		return
	}

	if u.Scheme != "https" && u.Scheme != "http" {
		err = errors.New("URL scheme must be 'http' or 'https'")
		return
	}

	if !strings.HasSuffix(u.Path, suffix) {
		err = errors.New("URL path must end with '.js' extension")
		return
	}

	u.RawQuery = ""
	u.Fragment = ""

	uSuffix, _ := url.Parse(u.String())
	uReplaceJs, _ := url.Parse(u.String())

	uSuffix.Path = u.Path + ".map"
	uReplaceJs.Path = u.Path[0:len(u.Path)-len(suffix)] + ".map"

	urls = []string{uSuffix.String(), uReplaceJs.String()}
	return
}
