package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	Count int
	Word  string
}

func Top10(inputText string) []string {
	words := strings.Fields(inputText)

	wordCountMap := make(map[string]int)

	for i := range words {
		wordCountMap[words[i]]++
	}

	wordCountList := make([]*wordCount, 0)
	for key, value := range wordCountMap {
		wordCount := &wordCount{
			Count: value,
			Word:  key,
		}
		wordCountList = append(wordCountList, wordCount)
	}

	sort.SliceStable(wordCountList, func(p, q int) bool {
		if wordCountList[p].Count == wordCountList[q].Count {
			return wordCountList[p].Word < wordCountList[q].Word
		}
		return wordCountList[p].Count > wordCountList[q].Count
	})

	resultWordList := make([]string, 0)
	for i := range wordCountList {
		if i > 9 {
			break
		}
		resultWordList = append(resultWordList, wordCountList[i].Word)
	}

	return resultWordList
}
