package main

import (
	"reflect"
	"testing"
)

func TestCharCount(t *testing.T) {
	tables := []struct {
		ss string
		mm map[string]int
	}{
		{"123LLLfff", map[string]int{
			"number": 3,
			"letter": 6,
		}},
		{"", map[string]int{}},
		{"s p a c e", map[string]int{
			"space":  4,
			"letter": 5,
		}},
		{"(words + 9876)", map[string]int{
			"space":       2,
			"punctuation": 2,
			"letter":      5,
			"symbol":      1,
			"number":      4,
		}},
	}

	for _, table := range tables {
		got := CharCount(table.ss)
		areEqual := reflect.DeepEqual(got, table.mm)
		if !areEqual {
			t.Errorf("CharCount was incorrect, given: %v, got: %v, want: %v", table.ss, got, table.mm)
		}
	}
}
