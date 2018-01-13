package main

import (
	"fmt"
)

func main() {
	var arr = []int{6, 5, 3, 1, 8, 7, 2, 4, 9}
	fmt.Println(arr)
	fmt.Println(MergeSort(arr))
	var test = make([]int, len(arr)*2, len(arr)*2)
	fmt.Println(test)
	test = append(test[:3], arr[7:]...)
	//test[7] = 1
	fmt.Println(test)
}

func MergeSort(arr []int) []int {
	var alen = len(arr)
	var ns = make([]int, alen, alen)

	var tmp = &arr
	var stmp = &ns
	for add := 1; add < alen; add *= 2 {
		fmt.Println(add)
		var start1 = 0
		var start2 = 0
		ik := 0
		for start1 = 0; start1 < alen; start1 = start1 + add {
			start2 = start1 + add
			var end1 = start1 + add
			var end2 = start2 + add
			for start1 < end1 && start2 < end2 && start1 < alen && start2 < alen {
				if (*tmp)[start1] < (*tmp)[start2] {
					(*stmp)[ik] = (*tmp)[start1]
					start1++
				} else {
					(*stmp)[ik] = (*tmp)[start2]
					start2++
				}
				ik++
				fmt.Println(*stmp)
			}

			for start1 < end1 && start1 < alen {
				(*stmp)[ik] = (*tmp)[start1]
				start1++
				ik++
				fmt.Println(*stmp)
			}

			for start2 < end2 && start2 < alen {
				(*stmp)[ik] = (*tmp)[start2]
				start2++
				ik++
				fmt.Println(*stmp)
			}
		}
		tmp, stmp = stmp, tmp
		for index, _ := range *stmp {
			(*stmp)[index] = 0
		}
	}
	return (*tmp)
}
