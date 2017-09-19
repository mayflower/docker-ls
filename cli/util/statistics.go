package util

import (
	"fmt"
	"os"

	"github.com/mayflower/docker-ls/lib/connector"
)

const STATISTICS_TEMPLATE = `Statistics:

API Requests:        %d
Token cache hits:
    API level:       %d
    Auth level:      %d
Token cache misses:
    API level:       %d
    Auth level:      %d
Token cache fails:
    API level:       %d
    Auth level:      %d
`

func DumpStatistics(statistics connector.Statistics) {
	fmt.Fprintf(os.Stderr, STATISTICS_TEMPLATE,
		statistics.Requests(),
		statistics.TokenCacheHitsAtApiLevel(),
		statistics.TokenCacheHitsAtAuthLevel(),
		statistics.TokenCacheMissesAtApiLevel(),
		statistics.TokenCacheMissesAtAuthLevel(),
		statistics.TokenCacheFailsAtApiLevel(),
		statistics.TokenCacheFailsAtAuthLevel(),
	)
}
