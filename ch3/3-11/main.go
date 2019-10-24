// page 74
// Comma prints its argument numbers with a comma at each power of 1000.
// no recursion, use bytes.Buffer
// handles floating-point numbers and a sign

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", Comma(os.Args[i]))
	}
}

func Comma(s string) string {
	var buf bytes.Buffer
	var sign bool
	var decimal string

	// if negative number
	if s[0] == '-' {
		sign = true
		s = s[1:]
	}

	if len(s) <= 3 {
		if sign {
			s = "-" + s
		}
		return s
	}

	// check for decimal number
	if strings.LastIndex(s, ".") != -1 {
		i := strings.LastIndex(s, ".")
		decimal = s[i:]
		s = s[:i]
	}

	// use modulus to determine number of ints before first comma
	pre := len(s) % 3

	// so we don't get ,123,456
	if pre == 0 {
		pre = 3
	}

	// add back the sign if needed
	if sign {
		buf.WriteByte('-')
	}
	// add from beginning of string to pre
	buf.WriteString(s[:pre])

	// for loop to add a comma and three numbers
	for i := pre; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	if len(decimal) != 0 {
		buf.WriteString(decimal)
	}

	return buf.String()
}
