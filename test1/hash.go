package main

import (
	"fmt"
	"math"
)

//1 lv func
func HashFuncTimes(num int, tableSize int) int {
	B := float64(num)
	A := 0.618
	na := float64(tableSize) * math.Mod(B*A, 1)
	fmt.Println(na)
	return int(na)
}

//2 lv func
func HashFunc(num int, ts int) int {
	return num % ts
}

type HashList struct {
	data int
	p    []HashList
}

func main() {
	var arr = []int{6, 5, 3, 1, 8, 7, 2, 4, 9}
	fmt.Println(arr)

	hl := make([]HashList, len(arr), len(arr))

	//demo
	var tmp1 = 37
	fmt.Println(HashFuncTimes(tmp1, len(arr)))
	fmt.Println(HashFunc(tmp1, len(arr)))

	ts := len(arr)
	//simple
	for _, num := range arr {
		k := HashFuncTimes(num, ts)
		if hl[k].data != 0 {
			hl[k].p = make([]HashList, ts, ts)
			k2 := HashFunc(num, ts)
			hl[k].p[k2].data = num
		} else {
			hl[k].data = num
		}
	}
	fmt.Println(hl)
}
