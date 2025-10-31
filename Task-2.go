package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wrdCnt(x string) map[string]int {
	x = strings.ToLower(x)
	mp := make(map[string]int)
	for i := range len(x) {
		if !unicode.IsLetter(rune(x[i])) {
			continue
		}
		_, ok := mp[string(x[i])]
		if !ok {
			mp[string(x[i])] = 1
		} else {
			mp[string(x[i])] += 1
		}
	}
	fmt.Println(mp)
	return mp
}

func Palind(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; j >= 0; i, j = i+1, j-1 {

		if s[i] != s[j] {
			return false
		}
	}
	return true

}

func main() {
	wrdCnt("hel-lo")
	fmt.Println(Palind("Ma-am"))
}
