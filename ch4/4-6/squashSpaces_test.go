package main

import (
	"reflect"
	"testing"
)

func TestSquashSpaces(t *testing.T) {
	tables := []struct {
		unsquashed []byte
		squashed   []byte
	}{
		{[]byte("so     many     spaces   "), []byte("so many spaces ")},
		{[]byte("nospaces"), []byte("nospaces")},
		{[]byte("  leading  spaces"), []byte(" leading spaces")},
		{[]byte("normal spaces"), []byte("normal spaces")},
	}

	for _, table := range tables {
		got := SquashSpaces(table.unsquashed)
		areEqual := reflect.DeepEqual(got, table.squashed)
		if !areEqual {
			t.Errorf("SquashSpaces was incorrect, \ngiven: %v, \ngot:   %v, \nwant:  %v", table.unsquashed, got, table.squashed)
		}
	}
}
