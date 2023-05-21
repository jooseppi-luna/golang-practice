package summed_pairs

import "sort"

type Pair struct {
    a, b interface{}
}

func FindSummedPairs(sum int, arr []int) []Pair {
	// sort array in-place
	sort.Ints(arr)

	// init the two indices
	front := 0
	back := len(arr) - 1

	var pairs []Pair
	for front < back {
		currSum := arr[front] + arr[back]
		if currSum == sum {
			pairs = append(pairs, Pair{
				arr[front], 
				arr[back],
			})
			front++
			back--
		} else if currSum > sum {
			back--
		} else {
			front++
		}
	}
	
	return pairs
}

