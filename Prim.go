package main

import (
	"fmt"
)

const MAX_SIZE int = 5

//为了看上去 好一些
const MAX_VALUE int = 9

func main() {
	fmt.Println("Prim")
	var gg Graph
	var vexs = []string{"A", "B", "C", "D", "E"}
	gg.vexnum = 5
	gg.vexs = vexs

	for i := 0; i < len(vexs); i++ {
		for j := 0; j < len(vexs); j++ {
			gg.matrix[i][j] = MAX_VALUE
		}
	}
	initGG(&gg)
	fmt.Println(gg.vexs)
	PrintG(gg, len(vexs))
}

func PrintG(gg Graph, l int) {
	for i := 0; i < l; i++ {
		fmt.Println(gg.matrix[i])
	}
}

type Graph struct {
	vexs   []string                //定点集合
	vexnum int                     //定点数量
	edgnum int                     //边数量
	matrix [MAX_SIZE][MAX_SIZE]int //邻接矩阵
}

func initGG(gg *Graph) {
	gg.matrix[0][1] = 5
	gg.matrix[0][2] = 3

	gg.matrix[1][0] = 5
	gg.matrix[1][3] = 7
	gg.matrix[1][4] = 4

	gg.matrix[2][0] = 3
	gg.matrix[2][3] = 6

	gg.matrix[3][1] = 7
	gg.matrix[3][2] = 6
	gg.matrix[3][4] = 1

	gg.matrix[4][1] = 4
	gg.matrix[4][3] = 1
}

func prim() {

}
