package largest-smallest-unsorted-arr

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLargestSmallestUnsortedArr(t *testing.T) {
	tests := []struct{
		name string
		arr []int
		largest int
		smallest int
	}{
		{
			name: "base case",
			arr: []int{7, 4, 24, 5, 17, 34, 21, 9, 4},
			largest: 34,
			smallest: 4,
		},
		{
			name: "one elem",
			arr: []int{7},
			largest: 7,
			smallest: 7,
		},
		{
			name: "all the same",
			arr: []int{7, 7, 7, 7, 7, 7, 7, 7},
			largest: 7,
			smallest: 7,
		},
		{
			name: "negative vals",
			arr: []int{7, 4, -24, 5, 17, 34, 21, -9, 0, 4},
			largest: 34,
			smallest: -24,
		},
		{
			name: "ascending",
			arr: []int{7, 8, 14, 15, 17, 24, 27, 29, 40},
			largest: 40,
			smallest: 7,
		},
		{
			name: "descending",
			arr: []int{97, 88, 74, 65, 57, 44, 37, 29, 20},
			largest: 97,
			smallest: 20,
		},
		{
			name: "solution at beginning",
			arr: []int{97, 8, 74, 65, 57, 44, 37, 29, 20},
			largest: 97,
			smallest: 8,
		},
		{
			name: "solution at end",
			arr: []int{74, 25, 57, 34, 47, 29, 50, 97, 8},
			largest: 97,
			smallest: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			largest, smallest := LargestSmallestUnsortedArr(tt.arr)
			assert.Equal(t, tt.largest, largest)
			assert.Equal(t, tt.smallest, smallest)
		})
	}
}

