package lib

import (
	"fmt"
)

const (
	APPLICATION_NAME = "docker-ls"
)

var staticVersion string = "not available --- use go generate and rebuild"
var dynamicVersion *string

var dynamicShortVersion *string

func Version() string {
	if dynamicVersion != nil {
		return *dynamicVersion
	} else {
		return staticVersion
	}
}

func ApplicationName() string {
	if dynamicShortVersion == nil {
		return APPLICATION_NAME
	}

	return fmt.Sprintf("%s/%s", APPLICATION_NAME, *dynamicShortVersion)
}

//go:generate go run ../generators/version.go -out dynamic_version.go -package lib
