package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r := 0, 0
	for k := 0; k < len(result); k++ {
		if l >= len(left) {
			result[k] = right[r]
			r++
		} else if r >= len(right) {
			result[k] = left[l]
			l++
		} else if left[l] < right[r] {
			result[k] = left[l]
			l++
		} else {
			result[k] = right[r]
			r++
		}
	}

	return result
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("unsorted array : ", arr)

	sortedArr := MergeSort(arr)
	fmt.Println("sortedArray : ", sortedArr)
}
