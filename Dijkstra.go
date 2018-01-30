package main

import (
	"fmt"
)

const MAX_SIZE int = 5
const MAX_VALUE int = 0

func main() {
	fmt.Println("Dijkstra")
	var gg Graph
	var vexs = []string{"B", "A", "C", "D", "E"}
	gg.vexnum = 5
	gg.vexs = vexs
	initGG(&gg, vexs)
	PrintG(gg, 5)
}

type Graph struct {
	vexs   []string                //定点集合
	vexnum int                     //定点数量
	edgnum int                     //边数量
	matrix [MAX_SIZE][MAX_SIZE]int //邻接矩阵
}

type Edge struct {
	start  string
	end    string
	weight int
}

func Dijkstra() {

}

func initGG(gg *Graph, vexs []string) {
	for i := 0; i < len(vexs); i++ {
		for j := 0; j < len(vexs); j++ {
			gg.matrix[i][j] = MAX_VALUE
		}
	}
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

	gg.edgnum = 12 / 2
}

func PrintG(gg Graph, l int) {
	for i := 0; i < l; i++ {
		fmt.Println(gg.matrix[i])
	}
}
