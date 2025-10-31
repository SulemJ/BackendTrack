package main

import (
	"fmt"
	"strings"
	"unicode"
)

func test(x string) map[string]int {
	x = strings.ToLower(x)
	mp := make(map[string]int)
	for i := range len(x) {
		if !unicode.IsLetter(rune(x[i])) {
			continue
		}
		val, ok := mp[string(x[i])]
		if !ok {
			mp[string(x[i])] = 1
		} else {
			mp[string(x[i])] += 1
		}
		fmt.Println(val, ok, mp)
	}
	return mp
	// fmt.Println(x)
}
func main() {
	str := "hel-lo"
	test(str)
}
