package lib

import (
	"testing"
)

func TestRefspecParse(t *testing.T) {
	ref := EmptyRefspec()
	err := ref.Set("foo:bar")

	if err != nil {
		t.Fatal(err)
	}

	if ref.Repository() != "foo" || ref.Reference() != "bar" {
		t.Fatal("reference failed to parse correctly")
	}
}

func TestRefspecParseMultiColons(t *testing.T) {
	ref := EmptyRefspec()
	err := ref.Set("foo:bar:baz")

	if err != nil {
		t.Fatal(err)
	}

	if ref.Repository() != "foo" || ref.Reference() != "bar:baz" {
		t.Fatal("reference failed to parse correctly")
	}
}

func TestRefspecParseInvalid(t *testing.T) {
	ref := EmptyRefspec()

	if ref.Set("foo") == nil {
		t.Fatal("references without colons should not parse")
	}
}
