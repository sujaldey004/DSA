package main

import "fmt"

func mergesort(arr []int, c chan []int) {
	if len(arr) < 2 {
		c <- arr
		return
	}
	mid := len(arr) / 2
	leftChan := make(chan []int)
	rightChan := make(chan []int)
	go mergesort(arr[:mid], leftChan)
	go mergesort(arr[mid:], rightChan)
	left := <-leftChan
	right := <-rightChan
	c <- merge(left, right)
}

func merge(left, right []int) []int {
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
	arr := []int{37, 12, 25, 8, 19, 45, 3, 29, 15}
	fmt.Println("Unsorted array : ", arr)
	c := make(chan []int)
	go mergesort(arr, c)
	sorted := <-c
	fmt.Println("Sorted array : ", sorted)
}
