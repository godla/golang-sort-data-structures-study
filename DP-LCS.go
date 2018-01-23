package main

import (
	"fmt"
)

func main() {
	//DP 是一种设计技术
	fmt.Println("Dynamic programming LCS")
	str1 := "ABCBDAB"
	str2 := "BDCABA"

	fmt.Println(str1[1] == 'B')

	l1 := len(str1)
	l2 := len(str2)
	fmt.Println(str1, str2, l1, l2)
	d2 := LCS(str1, str2, l1, l2)
	printLCS(d2, str1, l1, l2)
	fmt.Println("\nend")
}

func LCS(str1 string, str2 string, l1 int, l2 int) [10][10]int {

	var d1 [10][10]int //make([][]int, 10, 10)
	var d2 [10][10]int //make([][]int, 10, 10)

	for i := 0; i < l1; i++ {
		d1[i][0] = 0
	}
	for i := 0; i < l2; i++ {
		d1[0][i] = 0
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				d1[i][j] = d1[i-1][j-1] + 1
				d2[i][j] = 1
			} else if d1[i-1][j] >= d1[i][j-1] {
				d1[i][j] = d1[i-1][j]
				d2[i][j] = 3
			} else {
				d1[i][j] = d1[i][j-1]
				d2[i][j] = 2
			}
		}
	}
	for _, value := range d2 {
		fmt.Println(value)
	}
	fmt.Println("-------------------")
	for _, value := range d1 {
		fmt.Println(value)
	}
	return d2
}

func printLCS(b [10][10]int, str1 string, i int, j int) {

	if i < 0 || j < 0 {
		return
	}
	fmt.Println(i, j)
	if b[i][j] == 1 {
		printLCS(b, str1, i-1, j-1)
		fmt.Printf("%s", string(str1[i-1]))
	} else if b[i][j] == 3 {
		printLCS(b, str1, i-1, j)
	} else {
		printLCS(b, str1, i, j-1)
	}
}
