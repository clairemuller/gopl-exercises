// Page 8
// Prints command-line arguments.
// Modify to print index and value of each argument
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
	// same using range
	// for i, arg := range os.Args {
	// 	fmt.Println(i, arg)
	// }
}
