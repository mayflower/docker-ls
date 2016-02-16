package lib

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

var linkToNextHeaderRegexp *regexp.Regexp = regexp.MustCompile(`^<([^>]*)>;\s+rel="next"\s*$`)

func parseLinkToNextHeader(header string) (nextUrl *url.URL, err error) {
	match := linkToNextHeaderRegexp.FindAllStringSubmatch(header, -1)

	if len(match) != 1 {
		err = errors.New(fmt.Sprintf("malformed link header: %s", header))
		return
	}

	nextUrl, err = url.Parse(match[0][1])

	return
}
