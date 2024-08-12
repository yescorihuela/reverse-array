package main

import (
	"reflect"
	"testing"
)

func TestRotateArrays(t *testing.T) {
	testCases := []struct {
		name     string
		times    int
		array    []int
		expected []int
	}{
		{
			name:     "Once",
			times:    1,
			array:    []int{1, 2, 3, 4},
			expected: []int{4, 1, 2, 3},
		},
		{
			name:     "Twice",
			times:    2,
			array:    []int{1, 2, 3, 4},
			expected: []int{3, 4, 1, 2},
		},
		{
			name:     "Three times",
			times:    3,
			array:    []int{1, 2, 3, 4},
			expected: []int{2, 3, 4, 1},
		},
		{
			name:     "Four times",
			times:    4,
			array:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		out := rotateArray(tc.times, tc.array)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf("%s => test has failed", tc.name)
		}
	}
}
