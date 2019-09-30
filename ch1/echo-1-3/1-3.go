// Page 8
// Prints command-line arguments.
// Measure difference in running time
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	slowForLoop()
	fasterForLoop()
	fastestStringsJoin()
}

func slowForLoop() {
	start := time.Now()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	fmt.Println("runtime:", time.Since(start).Seconds())
}

func fasterForLoop() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
	fmt.Println("runtime:", time.Since(start).Seconds())
}

func fastestStringsJoin() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("runtime:", time.Since(start).Seconds())
}
