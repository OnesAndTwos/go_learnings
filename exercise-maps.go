package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

//WordCount - boo
func WordCount(s string) map[string]int {

	wordMap := make(map[string]int)

	for _, v := range strings.Fields(s) {
		wordMap[v]++
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
