// page 78
// write const declarations for KB, MB, up through YB as compactly as you can

package main

import (
	"fmt"
)

const (
	// kilobyte		2^10 or ~1,000 bytes
	KB = 1000
	// megabyte		2^20 or ~1,000,000 bytes
	MB = KB * 1000
	// gigabyte		2^30 or ~1,000,000,000 bytes
	GB = MB * 1000
	// terabyte		2^40 or ~1,000,000,000,000 bytes
	TB = GB * 1000
	// petabyte		2^50 or ~1,000,000,000,000,000 bytes
	PB = TB * 1000
	// exabyte		2^60 or ~1,000,000,000,000,000,000 bytes
	EB = PB * 1000
	// zettabyte	2^70 or ~1,000,000,000,000,000,000,000 bytes
	ZB = EB * 1000
	// yottabyte	2^80 or ~1,000,000,000,000,000,000,000,000 bytes
	YB = ZB * 1000
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB)
}
