package main

import (
	"fmt"
)

func main() {
	var arr = []int{6, 5, 3, 1, 8, 7, 2, 4, 9, 0, 3}
	fmt.Println(arr)
	CountingSort(arr)
}

func CountingSort(arr []int) {
	ln := len(arr)
	var sortAry = make([]int, ln, ln)
	var countAry = make([]int, ln, ln)

	for _, n := range arr {
		countAry[n]++
	}

	fmt.Println("count arr :", countAry)
	for i := 1; i < ln; i++ {
		countAry[i] += countAry[i-1]
	}
	fmt.Println("count arr add:", countAry)

	for _, kv := range arr {
		countAry[kv]--
		sortAry[countAry[kv]] = kv
	}

	fmt.Println(sortAry)
}
