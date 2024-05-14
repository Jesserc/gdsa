package main

import "fmt"

func main() {

	fmt.Println("Index of element =", binarySearch([]int{1, 2, 3, 4, 5}, 3))

}

// binarySearch implements a binary (divide & conquer) search algorithm which only works on sorted arrays.
func binarySearch(items []int, value int) int {
	low := 0
	high := len(items) - 1

	for low <= high {
		// fmt.Println("low", low) // uncomment this and similar ones to see the flow
		// fmt.Println("high", high)

		mid := (low + high) / 2
		// fmt.Println("mid", mid)

		if items[mid] == value {
			return mid
		}

		if items[mid] < value {
			low = mid + 1
			// fmt.Println("low", low)
		} else {
			high = mid - 1
			// fmt.Println("high", high)
		}
	}
	return -1
}
