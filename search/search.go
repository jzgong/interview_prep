package search

import "fmt"

func BinarySearchRecursive(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	mid := len(arr) / 2

	if arr[mid] == target {
		fmt.Println("inside")
		return true
	} else if arr[mid] < target {
		return BinarySearchRecursive(arr[mid+1:], target)
	} else {
		return BinarySearchRecursive(arr[:mid], target)
	}

	return false
}

func BinarySearch(arr []int, target int) bool {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (end + start) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}
