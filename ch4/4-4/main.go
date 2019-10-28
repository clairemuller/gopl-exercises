// pg 93
// implement a version of rotate that operates in a single pass

package main

import "fmt"

func main() {
	ss := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("before:", ss)
	Rotate(ss, 3)
	fmt.Println("after:", ss)

}

// Rotate rotates a slice to the left
func Rotate(ss []int, nn int) {
	for nn > 0 {
		temp := ss[0]
		copy(ss, ss[1:])
		ss[len(ss)-1] = temp
		nn--
	}
}
