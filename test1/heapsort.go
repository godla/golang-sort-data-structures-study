package main

import (
	"fmt"
)

func main() {
	array2 := []int{3, 2, 1}
	heapsort(array2)
	fmt.Println("HeapMY", array2)
}

func heapsort(array []int) {
	ep := (len(array) - 1) >> 1
	fmt.Println(ep)
	for i := ep; i >= 0; i-- {
		heapt(array, i, len(array)-1)
	}

	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapt(array, 0, i-1)
	}
}

func heapt(array []int, start int, end int) {
	le := start*2 + 1
	re := le + 1
	if le > end {
		return
	}

	var tmp = le
	if re <= end && array[re] > array[le] {
		tmp = re
	}

	if array[tmp] > array[start] {
		fmt.Println(start, end, array)
		array[start], array[tmp] = array[tmp], array[start]
		heapt(array, tmp, end)
	}
}
