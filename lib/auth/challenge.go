package auth

import (
	"fmt"
	"net/url"
	"regexp"
)

var challengeRegex *regexp.Regexp = regexp.MustCompile(
	`^\s*Bearer\s+realm="([^"]+)",service="([^"]+)",scope="([^"]+)"\s*$`)

var scopeSeparatorRegex *regexp.Regexp = regexp.MustCompile(`\s+`)

type Challenge struct {
	realm   *url.URL
	service string
	scope   []string
}

func ParseChallenge(challengeHeader string) (ch *Challenge, err error) {
	match := challengeRegex.FindAllStringSubmatch(challengeHeader, -1)

	if len(match) != 1 {
		err = fmt.Errorf("malformed challenge header: '%s'", challengeHeader)
	} else {
		var parsedRealm *url.URL
		parsedRealm, err = url.Parse(match[0][1])

		if err != nil {
			return
		}

		ch = &Challenge{
			realm:   parsedRealm,
			service: match[0][2],
			scope:   scopeSeparatorRegex.Split(match[0][3], -1),
		}
	}

	return
}

func (c *Challenge) buildRequestUrl() *url.URL {
	var authUrl url.URL = *c.realm
	var authParams url.Values = make(map[string][]string)

	authParams.Set("service", c.service)
	for _, scope := range c.scope {
		authParams.Add("scope", scope)
	}

	authUrl.RawQuery = authParams.Encode()

	return &authUrl
}
