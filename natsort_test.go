package natsort

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	// cases from https://github.com/skarademir/naturalsort/blob/master/naturalsort_test.go
	tests := []struct {
		data     []string
		expected []string
	}{
		{
			nil,
			nil,
		},
		{
			[]string{},
			[]string(nil), // modified case
		},
		{
			[]string{"a"},
			[]string{"a"},
		},
		{
			[]string{"0"},
			[]string{"0"},
		},
		{
			[]string{"data", "data20", "data3"},
			[]string{"data", "data3", "data20"},
		},
		{
			[]string{"1", "2", "30", "22", "0", "00", "3"},
			[]string{"00", "0", "1", "2", "3", "22", "30"}, // modified case
		},
		{
			[]string{"A1", "A0", "A21", "A11", "A111", "A2"},
			[]string{"A0", "A1", "A2", "A11", "A21", "A111"},
		},
		{
			[]string{"A1BA1", "A11AA1", "A2AB0", "B1AA1", "A1AA1"},
			[]string{"A1AA1", "A1BA1", "A2AB0", "A11AA1", "B1AA1"},
		},
		{
			[]string{"1ax10", "1a10", "1ax2", "1ax"},
			[]string{"1a10", "1ax", "1ax2", "1ax10"},
		},
		{
			[]string{"z1a10", "z1ax2", "z1ax"},
			[]string{"z1a10", "z1ax", "z1ax2"},
		},
		{
			// regression test for #8
			[]string{"a0001", "a0000001"},
			[]string{"a0000001", "a0001"}, // modified case
		},
		{
			// regression test for #10 - Number sort before any symbols even if theyre lower on the ASCII table
			[]string{"#1", "1", "_1", "a"},
			[]string{"1", "#1", "_1", "a"},
		},
		{
			// regression test for #10 - Number sort before any symbols even if theyre lower on the ASCII table
			[]string{"#1", "1", "_1", "a"},
			[]string{"1", "#1", "_1", "a"},
		},
		{ // test correct handling of space-only strings
			[]string{"1", " ", "0"},
			[]string{"0", "1", " "},
		},
		{ // test correct handling of multiple spaces being correctly ordered AFTER numbers
			[]string{"1", " ", " 1", "  "},
			[]string{"1", " ", " 1", "  "},
		},
		{
			[]string{"1", "#1", "a#", "a1"},
			[]string{"1", "#1", "a1", "a#"},
		},
		{
			// regression test for #10
			[]string{"111111111111111111112", "111111111111111111113", "1111111111111111111120"},
			[]string{"111111111111111111112", "111111111111111111113", "1111111111111111111120"},
		},
		{
			// regression test for #10
			[]string{"111111111111111111113", "1111111111111111111102", "1111111111111111111120", "111111111111111111112"},
			[]string{"111111111111111111112", "111111111111111111113", "1111111111111111111102", "1111111111111111111120"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestSort-%v", tt.data), func(t *testing.T) {
			sortedData := slices.SortedFunc(slices.Values(tt.data), Compare)
			if !reflect.DeepEqual(sortedData, tt.expected) {
				t.Errorf("Sort(%#v) = %#v, want %#v", tt.data, sortedData, tt.expected)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	// cases from https://github.com/maruel/natural/blob/main/natural_test.go
	tests := []struct {
		args []string
		want int
	}{
		{[]string{"", "a"}, -1},
		{[]string{"a", "b"}, -1},
		{[]string{"a", "aa"}, -1},
		{[]string{"a0", "a1"}, -1},
		{[]string{"a00", "a0"}, -1}, // modified case
		{[]string{"a00", "a01"}, -1},
		{[]string{"a01", "a1"}, -1},
		{[]string{"a01", "a2"}, -1},
		{[]string{"a01x", "a2x"}, -1},
		// Only the last number matters.
		{[]string{"a00b1", "a0b00"}, -1},  // modified case
		{[]string{"a00b01", "a0b00"}, -1}, // modified case
		{[]string{"a00b0", "a0b00"}, -1},
		{[]string{"a00b00", "a0b01"}, -1},
		{[]string{"a00b00", "a0b1"}, -1},
		{[]string{"a", ""}, 1},
		{[]string{"aa", "a"}, 1},
		{[]string{"b", "a"}, 1},
		{[]string{"a01", "a00"}, 1},
		{[]string{"a2", "a01"}, 1},
		{[]string{"a2x", "a01x"}, 1},
		{[]string{"a0b0", "a00b00"}, 1},  // modified case
		{[]string{"a0b00", "a00b01"}, 1}, // modified case
		{[]string{"10", "2"}, 1},
		{[]string{"a", "a"}, 0},
		{[]string{"a01", "a01"}, 0},
		{[]string{"a1", "a1"}, 0},
		// {[]string{"a00b00", "a0b00"}, 0}, // removed case
		{[]string{"a0b00", "a0b00"}, 0},
		// https://github.com/maruel/natural/issues/5
		{[]string{"a100000000000000000000a1", "a100000000000000000000a2"}, -1},
		{[]string{"a100000000000000000000a10", "a100000000000000000000a2"}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestCompare-%s-%s", tt.args[0], tt.args[1]), func(t *testing.T) {
			if got := Compare(tt.args[0], tt.args[1]); got != tt.want {
				t.Errorf("Compare(%s, %s) = %v, want %v", tt.args[0], tt.args[1], got, tt.want)
			}
		})
	}
}
