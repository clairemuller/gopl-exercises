// page 74
// given two strings, determine if they're anagrams of each other

package main

func main() {

}

func AreAnagrams(s1 string, s2 string) bool {
	// if the lengths aren't the same, can't be anagrams
	if len(s1) != len(s2) {
		return false
	}

	// make a map to count s1's letters
	hash := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		hash[s1[i]] += 1
	}

	// compare s2's letters to the map
	// if the letter doesn't exist, return false
	// else decrement the count
	for i := 0; i < len(s2); i++ {
		if hash[s2[i]] == 0 {
			return false
		} else {
			hash[s2[i]] -= 1
		}
	}

	return true
}
