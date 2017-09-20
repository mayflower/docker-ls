package tag

import (
	"encoding/json"
	"strings"

	"github.com/mayflower/docker-ls/cli/docker-ls/response"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
)

type Executor struct {
	CliConfig     *util.CliConfig
	LibraryConfig *lib.Config
	Tag           string
	RawManifest   bool
	ParseHistory  bool
}

func (e *Executor) Execute() (err error) {
	if e.CliConfig.InteractivePassword {
		err = util.PromptPassword(e.LibraryConfig)
		if err != nil {
			return
		}
	}

	ref := lib.EmptyRefspec()
	err = ref.Set(e.Tag)
	if err != nil {
		return
	}

	progress := util.NewProgressIndicator(e.CliConfig)
	progress.Start("requesting manifest")

	registryApi, err := lib.NewRegistryApi(*e.LibraryConfig)
	if err != nil {
		return
	}

	tagDetails, err := registryApi.GetTagDetails(ref, e.CliConfig.ManifestVersion)

	progress.Progress()
	progress.Finish("done")

	if err != nil {
		return
	}

	if e.RawManifest {
		manifest := tagDetails.RawManifest()

		if e.ParseHistory {
			e.parseHistory(manifest)
		}

		err = util.SerializeToStdout(manifest, e.CliConfig)
	} else {
		err = util.SerializeToStdout(response.NewTagDetailsL0(tagDetails, true), e.CliConfig)
	}

	return
}

func (e *Executor) parseHistory(rawManifest interface{}) {
	var manifest map[string]interface{}
	var ok bool

	if manifest, ok = rawManifest.(map[string]interface{}); !ok {
		return
	}

	var history []interface{}

	if _, ok = manifest["history"]; !ok {
		return
	}
	if history, ok = manifest["history"].([]interface{}); !ok {
		return
	}

	for _, rawItem := range history {
		var item map[string]interface{}

		if item, ok = rawItem.(map[string]interface{}); !ok {
			continue
		}
		if _, ok = item["v1Compatibility"]; !ok {
			continue
		}

		var rawHistoryItem string

		if rawHistoryItem, ok = item["v1Compatibility"].(string); !ok {
			continue
		}

		decoder := json.NewDecoder(strings.NewReader(rawHistoryItem))
		var decodedHistoryItem interface{}

		if decoder.Decode(&decodedHistoryItem) != nil {
			continue
		}

		item["v1Compatibility"] = decodedHistoryItem
	}
}
