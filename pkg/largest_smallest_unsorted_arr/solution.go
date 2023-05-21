package largest_smallest_unsorted_arr

func LargestSmallestUnsortedArr(arr []int) (int, int) {
	largest := arr[0]
	smallest := arr[0]
	for _, elem := range arr {
		if largest < elem {
			largest = elem
		}
		if smallest > elem {
			smallest = elem
		}
	}

	return largest, smallest
}

