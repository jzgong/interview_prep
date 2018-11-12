package sorting

import "golang.org/x/exp/rand"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivot := rand.Int() % len(arr)

	// Move the pivot to the right
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Pile elements smaller than the pivot
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smallest element
	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr

}

func RadixStr(arr []string, index int) []string {
	if len(arr) <= 1 {
		return arr
	}

	doneBucket := []string{}
	var bucketList [27][]string

	// Go over arr and put strings into buckets by letter
	for i := range arr {
		if len(arr[i]) <= index {
			doneBucket = append(doneBucket, arr[i])
		} else {
			asciiIndex := arr[i][index] - "a"[0]
			bucketList[asciiIndex] = append(bucketList[asciiIndex], arr[i])
		}
	}

	// Recursively apply to smaller buckets
	var tempBucketList [][]string
	for _, bucket := range bucketList {
		tempBucketList = append(tempBucketList, RadixStr(bucket, index+1))
	}

	// Add sorted items to doneBucket
	for _, bucket := range tempBucketList {
		doneBucket = append(doneBucket, bucket...)
	}

	return doneBucket
}

func RadixInt(arr []int) []int {
	max_val := 0

	for i := range arr {
		if max_val < arr[i] {
			max_val = arr[i]
		}
	}

	digit := 1

	for max_val/digit > 0 {
		var bucketList [10][]int

		for i := range arr {
			bucketList[(arr[i]/digit)%10] = append(bucketList[(arr[i]/digit)%10], arr[i])
		}

		arr = []int{}

		for _, bucket := range bucketList {
			arr = append(arr, bucket...)
		}

		digit *= 10
	}

	return arr
}
