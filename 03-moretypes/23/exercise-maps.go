package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount : I made this code!
func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	ret := make(map[string]int)
	for _, value := range fields {
		ret[value]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
