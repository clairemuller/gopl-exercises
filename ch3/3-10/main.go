// page 74
// Comma prints its argument numbers with a comma at each power of 1000.
// no recursion, use bytes.Buffer

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer

	if len(s) <= 3 {
		return s
	}

	// use modulus to determine number of ints before first comma
	pre := len(s) % 3

	// so we don't get ,123,456
	if pre == 0 {
		pre = 3
	}

	// add from beginning of string to pre
	buf.WriteString(s[:pre])

	// for loop to add a comma and three numbers
	for i := pre; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
