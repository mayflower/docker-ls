package lib

var staticVersion string = "not available"
var dynamicVersion *string

func Version() string {
	if dynamicVersion != nil {
		return *dynamicVersion
	} else {
		return staticVersion
	}
}

//go:generate go run ../generators/version.go -out dynamic_version.go -package lib
