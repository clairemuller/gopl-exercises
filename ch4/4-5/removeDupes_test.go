package main

import (
	"reflect"
	"testing"
)

func TestRemoveDupes(t *testing.T) {
	tables := []struct {
		ss  []string
		ss2 []string
	}{
		{[]string{"abc", "abc", "abc", "abc", "abc"}, []string{"abc"}},
		{[]string{"abc", "wow", "abc"}, []string{"abc", "wow", "abc"}},
		{[]string{"abc", "wow", "abc", "wow"}, []string{"abc", "wow", "abc", "wow"}},
		{[]string{}, []string{}},
		{[]string{""}, []string{""}},
		{[]string{"", "", "", ""}, []string{""}},
		{[]string{"abc", "ab", "ab"}, []string{"abc", "ab"}},
		{[]string{"abc", "abc", "ab"}, []string{"abc", "ab"}},
	}

	for _, table := range tables {
		got := RemoveDupes(table.ss)
		areEqual := reflect.DeepEqual(got, table.ss2)
		if !areEqual {
			t.Errorf("RemoveDupes was incorrect, given: %v, got: %v, want: %v", table.ss, got, table.ss2)
		}
	}
}
