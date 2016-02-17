package main

import (
	"fmt"
	"os"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
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

func dumpStatistics(statistics lib.Statistics) {
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
