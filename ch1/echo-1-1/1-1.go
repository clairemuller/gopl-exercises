// Page 8
// Prints command-line arguments.
// Modify to also print os.Args[0], the command that invoked it
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])
}
