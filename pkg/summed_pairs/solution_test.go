package summed_pairs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindSummedPairs(t *testing.T) {
	tests := []struct{
		name string
		arr []int
		sum int
		pairs []Pair
	}{
		{
			name: "base case",
			arr: []int{8, 3, 12, 9, 15, 14, 11, 5, 6, 1},
			sum: 15,
			pairs: []Pair{Pair{1, 14}, Pair{3, 12}, Pair{6, 9}},
		},
		{
			name: "one elem",
			arr: []int{7},
			sum: 7,
			pairs: []Pair(nil),
		},
		{
			name: "all the same, don't add up",
			arr: []int{7, 7, 7, 7, 7, 7, 7, 7},
			sum: 13,
			pairs: []Pair(nil),
		},
		{
			name: "all the same, do add up",
			arr: []int{6, 6, 6, 6, 6, 6, 6, 6},
			sum: 12,
			pairs: []Pair{Pair{6, 6}, Pair{6, 6}, Pair{6, 6}, Pair{6, 6}},
		},
		{
			name: "negative vals",
			arr: []int{-22, -18, -13, -8, -7, 0, 3, 5, 7, 9, 11, 12, 19, 33, 34},
			sum: 12,
			pairs: []Pair{Pair{-22, 34}, Pair{-7, 19}, Pair{0, 12}, Pair{3, 9}, Pair{5, 7}},
		},
		{
			name: "all too small",
			arr: []int{1, 4, 7, 8, 10, 11, 13, 27},
			sum: 43,
			pairs: []Pair(nil),
		},
		{
			name: "don't add up",
			arr: []int{1, 4, 7, 8, 10, 11, 13, 27},
			sum: 13,
			pairs: []Pair(nil),
		},
		{
			name: "all too big",
			arr: []int{1, 4, 7, 8, 10, 11, 13, 27},
			sum: 1,
			pairs: []Pair(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs := FindSummedPairs(tt.sum, tt.arr)
			assert.Equal(t, tt.pairs, pairs)
		})
	}
}

