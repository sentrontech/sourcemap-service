package sourcemap

import (
	"net/url"
	"regexp"
)

// regex "borrowed" from https://github.com/lydell/source-map-url/blob/master/source-map-url.js#L16..L27
const pattern = `(?:\/\*(?:\s*\r?\n(?:\/\/)?)?(?:[#@] sourceMappingURL=([^\s'"]*))\s*\*\/|\/\/(?:[#@] sourceMappingURL=([^\s'"]*)))\s*`

var re = regexp.MustCompile(pattern)

// ParseURL returns the source map URL from a file, if one exists
func ParseURL(jsFileURL string, fileContents string) (string, error) {
	match := re.FindStringSubmatch(fileContents)
	if len(match) < 2 {
		return "", nil
	}
	rel := match[2]

	u, err := url.Parse(rel)
	if err != nil {
		return "", err
	}
	base, err := url.Parse(jsFileURL)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(u).String(), nil
}
