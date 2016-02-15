package auth

import (
	"net/url"
	"reflect"
	"testing"
)

func testParse(t *testing.T, header string, fixture Challenge) {
	challenge, err := ParseChallenge(header)

	if err != nil {
		return
	}

	if expected, actual := fixture.realm.String(), challenge.realm.String(); expected != actual {
		t.Fatalf("realm failed to parse; got %s, expected %s", actual, expected)
	}

	if expected, actual := fixture.service, challenge.service; expected != actual {
		t.Fatalf("service failed to parse; got %s, expected %s", actual, expected)
	}

	if expected, actual := fixture.scope, challenge.scope; !reflect.DeepEqual(expected, actual) {

		t.Fatalf("challenge failed to parse; got %v, expected %v", actual, expected)
	}
}

func testAuthUrl(t *testing.T, header, scheme, host, path string, scopes []string) {
	challenge, err := ParseChallenge(header)

	if err != nil {
		t.Fatal(err)
	}

	authUrl := challenge.buildRequestUrl()

	if expected, actual := scheme, authUrl.Scheme; expected != actual {
		t.Fatalf("auth URL scheme not constructed correctly; got %s, expected %s", actual, expected)
	}

	if expected, actual := host, authUrl.Host; expected != actual {
		t.Fatalf("auth URL host not constructed correctly; got %s, expected %s", actual, expected)
	}

	if expected, actual := path, authUrl.Path; expected != actual {
		t.Fatalf("auth URL path not constructed correctly; got %s, expected %s", actual, expected)
	}

	if expected, actual := scopes, authUrl.Query()["scope"]; !reflect.DeepEqual(expected, actual) {
		t.Fatalf("auth URL schemes not constructed correctly; got %v, expected %v", actual, expected)
	}
}

func TestSimple(t *testing.T) {
	testcase := `Bearer realm="https://auth.docker.io/token",service="registry.docker.io",scope="repository:samalba/my-app:pull,push"`

	fixtureUrl, _ := url.Parse("https://auth.docker.io/token")
	testParse(t, testcase, Challenge{
		realm:   fixtureUrl,
		service: "registry.docker.io",
		scope: []string{
			"repository:samalba/my-app:pull,push",
		},
	})

	testAuthUrl(t, testcase, "https", "auth.docker.io", "/token", []string{
		"repository:samalba/my-app:pull,push",
	})
}

func TestHeaderPort(t *testing.T) {
	testcase := `Bearer realm="https://auth.docker.io:8888/token",service="registry.docker.io:9999",scope="repository:samalba/my-app:pull,push"`

	fixtureUrl, _ := url.Parse("https://auth.docker.io:8888/token")
	testParse(t, testcase, Challenge{
		realm:   fixtureUrl,
		service: "registry.docker.io:9999",
		scope: []string{
			"repository:samalba/my-app:pull,push",
		},
	})

	testAuthUrl(t, testcase, "https", "auth.docker.io:8888", "/token", []string{
		"repository:samalba/my-app:pull,push",
	})
}

func TestParseHeaderIpPort(t *testing.T) {
	testcase := `Bearer realm="https://192.168.1.1:8888/token",service="192.168.1.2:8888",scope="repository:samalba/my-app:pull,push"`

	fixtureUrl, _ := url.Parse("https://192.168.1.1:8888/token")
	testParse(t, testcase, Challenge{
		realm:   fixtureUrl,
		service: "192.168.1.2:8888",
		scope: []string{
			"repository:samalba/my-app:pull,push",
		},
	})

	testAuthUrl(t, testcase, "https", "192.168.1.1:8888", "/token", []string{
		"repository:samalba/my-app:pull,push",
	})
}

func TestParseHeaderMultiScope(t *testing.T) {
	testcase := `Bearer realm="https://auth.docker.io/token",service="registry.docker.io",scope="repository:samalba/my-app:pull,push repository:samalba/my-app:foo,bar"`

	fixtureUrl, _ := url.Parse("https://auth.docker.io/token")
	testParse(t, testcase, Challenge{
		realm:   fixtureUrl,
		service: "registry.docker.io",
		scope: []string{
			"repository:samalba/my-app:pull,push",
			"repository:samalba/my-app:foo,bar",
		},
	})

	testAuthUrl(t, testcase, "https", "auth.docker.io", "/token", []string{
		"repository:samalba/my-app:pull,push",
		"repository:samalba/my-app:foo,bar",
	})
}

func TestInvalid(t *testing.T) {
	_, err := ParseChallenge(`Bearrer realm="https://auth.docker.io/token",service="registry.docker.io",scope="repository:samalba/my-app:pull,push"`)

	if err == nil {
		t.Fatal("parsing an invalid challenge header should fail")
	}
}
