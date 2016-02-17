package lib

import (
	"reflect"
	"testing"
)

func TestLinkHeaderParse(t *testing.T) {
	testcase := `<http://example.com/v2/_catalog?n=20&last=b>; rel="next"`

	url, err := parseLinkToNextHeader(testcase)

	if err != nil {
		t.Fatal(err)
	}

	if url.Scheme != "http" ||
		url.Host != "example.com" ||
		url.Path != "/v2/_catalog" ||
		!reflect.DeepEqual((map[string][]string)(url.Query()), map[string][]string{
			"n":    []string{"20"},
			"last": []string{"b"},
		}) {

		t.Fatalf("parsing failed; got %s", url.String())
	}
}

func TestLinkHeaderParseInvalid(t *testing.T) {
	testcase := `<http://example.com/v2/_catalog?n=20&last=b; rel="next"`

	_, err := parseLinkToNextHeader(testcase)

	if err == nil {
		t.Fatal("parsing an invalid header should fail")
	}
}
