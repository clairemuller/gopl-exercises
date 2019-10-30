// page 93
// write an in-place function that squashes each run of adjacent Unicode spaces
// in a UTF-8 encoded []byte slice into a single ASCII space

package main

import (
	"unicode"
)

func main() {
}

// SquashSpaces starts at index 0
// if the element is a space, check if the next element is a space
// if so, shift everything to the left, copying over the extra space
// continue
func SquashSpaces(bb []byte) []byte {
	i := 0
	for i < len(bb) {
		if unicode.IsSpace(rune(bb[i])) && (i != len(bb)-1) {
			if unicode.IsSpace(rune(bb[i+1])) {
				copy(bb[i:], bb[i+1:])
				bb = bb[:len(bb)-1]
				continue
			}
		}
		i++
	}
	return bb
}
