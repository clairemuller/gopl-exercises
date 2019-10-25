package main

import "testing"

func TestComma(t *testing.T) {
	tables := []struct {
		num      string
		commaNum string
	}{
		{"0", "0"},
		{"-1", "-1"},
		{"12", "12"},
		{"1.2", "1.2"},
		{"123", "123"},
		{"-123", "-123"},
		{"1234", "1,234"},
		{"-1234", "-1,234"},
		{"12345", "12,345"},
		{"-12345", "-12,345"},
		{"123456", "123,456"},
		{"-123456", "-123,456"},
		{"1234567", "1,234,567"},
		{"-1234567", "-1,234,567"},
		{"1.234", "1.234"},
		{"-1234.5", "-1,234.5"},
		{"-12345.67", "-12,345.67"},
	}

	for _, table := range tables {
		n := Comma(table.num)
		if n != table.commaNum {
			t.Errorf("Comma was incorrect, given: %v, got: %v, want: %v", table.num, n, table.commaNum)
		}
	}
}
