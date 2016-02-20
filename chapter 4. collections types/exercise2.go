package main

import (
	"golang.org/x/tour/wc"	
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}
	return count
	
}

func main() {
	wc.Test(WordCount)
}
