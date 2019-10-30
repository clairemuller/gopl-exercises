// pg 93
// rewrite reverse to use an array pointer instead of a slice

package main

import "fmt"

func main() {
	aa := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("before:", aa)
	reverse(&aa)
	fmt.Println("after:", aa)

}

func reverse(aa *[6]int) {
	for i, j := 0, len(aa)-1; i < j; i, j = i+1, j-1 {
		aa[i], aa[j] = aa[j], aa[i]
	}
}
