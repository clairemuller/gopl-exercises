// pg 93
// write an in-place function to eliminate adjacent duplicates in a []string slice

package main

func main() {

}

// RemoveDupes starts at index 1, checking it against previous element
// if they're the same, copy over the element by sliding everything to the left once
// and removing the last element by reassigning the slice
// otherwise, move to next element
func RemoveDupes(ss []string) []string {
	i := 1
	for i < len(ss) {
		if ss[i] == ss[i-1] {
			copy(ss[i:], ss[i+1:])
			ss = ss[:len(ss)-1]
		} else {
			i++
		}
	}
	return ss
}
