package main

import "cmp"

// mergeSortedSlices merges two sorted slices into one sorted slice.
//	It takes two sorted slices as input parameters and returns a new sorted slice containing all elements from both input slices.
//	If either input slice is empty, it returns the other slice.
//	The function uses a merge sort algorithm to combine the slices efficiently.
/*

	Example:
	mergeSortedSlices([]int{1, 3, 5, 7, 20, 30, 40}, []int{2, 4, 6, 50})
	len(slice1) == 7, len(slice2) == 4

	Loop steps (i and j values on each step):
	i :=	1, 2, 3, 4, 5, 6, 7
	j :=	0, 1, 2, 3

	Sorted slice: {1, 2, 3, 4, 5, 6, 7, 20, 30, 40}
*/
func mergeSortedSlices[T cmp.Ordered](slice1, slice2 []T) []T {
	// If either slice is empty, return the other slice
	if len(slice1) == 0 {
		return slice2
	}

	if len(slice2) == 0 {
		return slice1
	}

	// Initialize an empty slice to store the merged result
	sorted := make([]T, 0, len(slice1)+len(slice2))

	// Initialize indexes for both slices
	i := 0
	j := 0

	// Loop until we have processed all elements from either slice.
	// This loop proceeds only when `i` and `j` are less than the length of their respective slice.
	// Note that the incrementation of either `i` or `j` does not happen linearly because of the if-else block.
	for i < len(slice1) && j < len(slice2) {
		// Compare elements from both slices and append the smaller one to the sorted slice
		if slice1[i] < slice2[j] {
			sorted = append(sorted, slice1[i])
			i++
		} else {
			sorted = append(sorted, slice2[j])
			j++
		}
	}

	// Append any remaining elements from slice1
	for i < len(slice1) {
		sorted = append(sorted, slice1[i])
		i++
	}

	// Append any remaining elements from slice2
	for j < len(slice2) {
		sorted = append(sorted, slice2[j])
		j++
	}

	// Return the merged and sorted slice
	return sorted
}
