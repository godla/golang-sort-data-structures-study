package main

import (
	"fmt"
)

//为了看上去 好一些
const MAX_VALUE int = 9

func main() {
	var arr = []int{6, 5, 5, 3, 1, 8, 7, 2, 4, 9}
	start := 0
	end := len(arr) - 1
	fmt.Println(arr)
	QuickSortRecursive(arr, start, end)
}

func QuickSortRecursive(arr []int, start int, end int) {
	if start >= end {
		return
	}
	k := arr[end]
	left := start
	right := end - 1
	for left < right {
		for arr[left] < k && left < right {
			left++
		}
		for arr[right] >= k && left < right {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
		fmt.Println(arr)
	}
	if arr[left] < arr[end] {
		left++
	}
	arr[left], arr[end] = k, arr[left]
	fmt.Println(arr)
	QuickSortRecursive(arr, start, left-1)
	QuickSortRecursive(arr, left+1, end)
}
