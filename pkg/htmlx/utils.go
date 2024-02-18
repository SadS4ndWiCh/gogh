package htmlx

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func stringIncludesWords(target string, words string) bool {
	if target == words {
		return true
	}

	targetSet := mapset.NewSet[string]()
	for _, s := range strings.Split(strings.TrimSpace(target), " ") {
		targetSet.Add(s)
	}

	wordsSet := mapset.NewSet[string]()
	for _, s := range strings.Split(strings.TrimSpace(words), " ") {
		wordsSet.Add(s)
	}

	return wordsSet.IsSubset(targetSet)
}
