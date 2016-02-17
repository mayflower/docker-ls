package lib

func cacheHintRegistryList() string {
	return "catalog:"
}

func cacheHintTagList(repository string) string {
	return "pull:" + repository
}

func cacheHintTagDetails(repository string) string {
	return "pull:" + repository
}
