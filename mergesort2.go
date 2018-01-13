package main

import (
	"fmt"
)

func main() {
	var arr = []int{6, 5, 3, 1, 8, 7, 2, 4, 9}
	var tmp = make([]int, len(arr), len(arr))
	start := 0
	end := len(arr)
	fmt.Println(arr)
	MergeSortRecursive(arr, tmp, start, end-1)
	fmt.Printf("finish :%v \n", arr)
	fmt.Printf("finish :%v \n", tmp)
}

func MergeSortRecursive(arr []int, tmp []int, start int, end int) {
	fmt.Printf("run : %d-%d\n", start, end)
	if start >= end {
		return
	}
	fmt.Printf("runing : %d-%d\n", start, end)
	var start1 int = start
	var end1 int = (end-start)>>1 + start1

	var start2 int = end1 + 1
	var end2 int = end

	var k = start1
	MergeSortRecursive(arr, tmp, start1, end1)
	MergeSortRecursive(arr, tmp, start2, end2)

	for ; start1 <= end1 && start2 <= end2; k++ {
		if arr[start1] < arr[start2] {
			tmp[k] = arr[start1]
			start1++
		} else {
			tmp[k] = arr[start2]
			start2++
		}
		fmt.Println(tmp)
	}
	for ; start1 <= end1; k++ {
		tmp[k] = arr[start1]
		start1++
		fmt.Println(tmp)
	}

	for ; start2 <= end2; k++ {
		tmp[k] = arr[start2]
		start2++
		fmt.Println(tmp)
	}

	for i := start; i <= end; i++ {
		arr[i] = tmp[i]
		tmp[i] = 0
		// unsort := append([]int{}, arr[end:]...)
		// tmp2 := tmp[i:end]
		// arr = append(append(arr[i:], tmp2[0:]...), unsort[0:]...)
	}
	fmt.Printf("last arr: %v\n", arr)
	return
}
