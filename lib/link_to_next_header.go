package lib

import (
	"fmt"
	"net/url"
	"regexp"
)

var linkToNextHeaderRegexp *regexp.Regexp = regexp.MustCompile(`^<([^>]*)>;\s+rel="next"\s*$`)

func parseLinkToNextHeader(header string) (nextUrl *url.URL, err error) {
	match := linkToNextHeaderRegexp.FindAllStringSubmatch(header, -1)

	if len(match) != 1 {
		err = fmt.Errorf("malformed link header: %s", header)
		return
	}

	nextUrl, err = url.Parse(match[0][1])

	return
}
