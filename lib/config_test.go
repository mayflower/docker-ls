package lib

import (
	"fmt"
	"os"
	"testing"
)

func TestInitRegistryURL(t *testing.T) {
	initRegistryURL()

	if DEFAULT_REGISTRY_URL_STRING != "https://index.docker.io" {
		t.Fatal("Expected DEFAULT_REGISTRY_URL_STRING to be \"https://docker.mycompany.com\"")
	}
}

func TestInitRegistryURLWithEnvVar(t *testing.T) {
	customRegistryURLString := "https://docker.mycompany.com"
	os.Setenv("DOCKER_REGISTRY_URL", customRegistryURLString)
	initRegistryURL()
	os.Unsetenv("DOCKER_REGISTRY_URL")

	if DEFAULT_REGISTRY_URL_STRING != customRegistryURLString {
		t.Fatal(fmt.Sprintf("Expected DEFAULT_REGISTRY_URL_STRING to be \"%s\"", customRegistryURLString))
	}
}
