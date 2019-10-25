package main

import "testing"

func TestAreAnagrams(t *testing.T) {
	tables := []struct {
		s1  string
		s2  string
		boo bool
	}{
		{"anagram", "anagram", true},
		{"anagram", "ranagam", true},
		{"anagram", "nope", false},
		{"", "", true},
		{"./", ".", false},
		{"0", "0", true},
		{"123", "0", false},
		{"anagram", "anagrat", false},
		{"anagram", "rats", false},
		{"abc", "abc", true},
		{"abc", "abcd", false},
	}

	for _, table := range tables {
		boolean := AreAnagrams(table.s1, table.s2)
		if boolean != table.boo {
			t.Errorf("AreAnagrams was incorrect, given: %v and %v, got: %v, want: %v", table.s1, table.s2, boolean, table.boo)
		}
	}
}
