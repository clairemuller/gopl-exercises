// pg 97
// modify charcount to count letters, digits, and so on in their Unicode categories

package main

import (
	"unicode"
)

func main() {
}

// CharCount takes a string and counts occurences of Unicode categories
func CharCount(ss string) map[string]int {
	mm := make(map[string]int)
	for _, rr := range ss {
		if unicode.IsLetter(rr) {
			mm["letter"]++
		} else if unicode.IsNumber(rr) {
			mm["number"]++
		} else if unicode.IsSpace(rr) {
			mm["space"]++
		} else if unicode.IsControl(rr) {
			mm["control"]++
		} else if unicode.IsSymbol(rr) {
			mm["symbol"]++
		} else if unicode.IsPunct(rr) {
			mm["punctuation"]++
		}
	}
	return mm
}
