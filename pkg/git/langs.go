package git

import (
	"fmt"
	"sort"
)

// SortLanguagesWithInts sorts languages from the given map to a slice where the first one is with the highest number
func SortLanguagesWithInts(langsWithSizes map[string]int) []string {
	var contentSizes []string
	reversedMap := map[string]string{}
	for lang, size := range langsWithSizes {
		key := fmt.Sprintf("%d_%s", size, lang)
		reversedMap[key] = lang
		contentSizes = append(contentSizes, key)
	}
	return sortLanguages(contentSizes, reversedMap)
}

// SortLanguagesWithFloats32 sorts languages from the given map to a slice where the first one is with the highest number
func SortLanguagesWithFloats32(langsWithSizes map[string]float32) []string {
	var contentSizes []string
	reversedMap := map[string]string{}
	for lang, size := range langsWithSizes {
		key := fmt.Sprintf("%f_%s", size, lang)
		reversedMap[key] = lang
		contentSizes = append(contentSizes, key)
	}
	return sortLanguages(contentSizes, reversedMap)
}

func sortLanguages(contentSizes []string, reversedMap map[string]string) []string {
	sort.Strings(contentSizes)

	var sortedLangs []string
	for i := len(contentSizes) - 1; i >= 0; i-- {
		sortedLangs = append(sortedLangs, reversedMap[contentSizes[i]])
	}
	return sortedLangs
}