package util

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	CLI_OPTION_RECURSION_LEVEL = 1 << iota
	CLI_OPTION_STATISTICS
	CLI_OPTION_PROGRESS
	CLI_OPTION_JSON_OUTPUT
	CLI_OPTION_MANIFEST_VERSION
	CLI_OPTION_INTERACTIVE_PASSWORD
	CLI_OPTION_TABLE_OUTPUT
)

const (
	CLI_OPTIONS_FULL = 0xFFFF
	CLI_OPTIONS_NONE = 0
)

type CliConfig struct {
	RecursionLevel      uint
	ManifestVersion     uint
	Statistics          bool
	Progress            bool
	JsonOutput          bool
	InteractivePassword bool
	TableOutput         bool
}

func AddCliConfigToFlags(flags *pflag.FlagSet, options uint) {
	c := NewCliConfig()

	if options&CLI_OPTION_RECURSION_LEVEL != 0 {
		flags.UintP("recursion", "l", c.RecursionLevel, "level of recursion")
	}

	if options&CLI_OPTION_STATISTICS != 0 {
		flags.Bool("statistics", c.Statistics, "show connection statistics")
	}

	if options&CLI_OPTION_PROGRESS != 0 {
		flags.Bool("progress-indicator", c.Progress, "show progress indicator")
	}

	if options&CLI_OPTION_JSON_OUTPUT != 0 {
		flags.BoolP("json", "j", c.JsonOutput, "output JSON instead of YAML")
	}

	if options&CLI_OPTION_MANIFEST_VERSION != 0 {
		flags.Uint("manifest-version", c.ManifestVersion, "manifest version to request")
	}

	if options&CLI_OPTION_INTERACTIVE_PASSWORD != 0 {
		flags.BoolP("interactive-password", "i", c.InteractivePassword, "prompt for password")
	}

	if options&CLI_OPTION_TABLE_OUTPUT != 0 {
		flags.Bool("table", c.TableOutput, "output table instead of YAML")
	}
}

func CliConfigFromViper() *CliConfig {
	return &CliConfig{
		RecursionLevel:      uint(viper.GetInt("recursion")),
		Statistics:          viper.GetBool("statistics"),
		Progress:            viper.GetBool("progress-indicator"),
		JsonOutput:          viper.GetBool("json"),
		ManifestVersion:     uint(viper.GetInt("manifest-version")),
		InteractivePassword: viper.GetBool("interactive-password"),
		TableOutput:         viper.GetBool("table"),
	}
}

func NewCliConfig() *CliConfig {
	return &CliConfig{
		RecursionLevel:      0,
		Statistics:          false,
		Progress:            true,
		JsonOutput:          false,
		ManifestVersion:     2,
		InteractivePassword: false,
		TableOutput:         false,
	}
}
