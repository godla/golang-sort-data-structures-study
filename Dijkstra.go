package main

import (
	"fmt"
)

const MAX_SIZE int = 5
const MAX_VALUE int = 9999

func main() {
	fmt.Println("Dijkstra")
	var gg Graph
	var vexs = []string{"A", "B", "C", "D", "E"}
	gg.vexnum = 5
	gg.vexs = vexs
	initGG(&gg, vexs)
	PrintG(gg, 5)
	Dijkstra(&gg, 1)
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

func Dijkstra(gg *Graph, start int) {

	var dist [MAX_SIZE]int //路劲长度数组
	var flag [MAX_SIZE]bool
	var prev [MAX_SIZE]int

	//init
	dist = gg.matrix[start]
	flag[start] = true //find start to start
	dist[start] = 0    //start to start length

	k := 0
	//广度搜索
	for i := 0; i < gg.vexnum; i++ {
		min := MAX_VALUE
		//find min
		for j := 0; j < gg.vexnum; j++ {
			if flag[j] == true && dist[j] < min {
				min = dist[j]
				k = j
			}
		}

		//set find
		flag[k] = true

		//update dist length
		for u := 0; u < gg.vexnum; u++ {
			weigth := 0
			if gg.matrix[k][u] == MAX_VALUE {
				weigth = MAX_VALUE
			} else {
				weigth = gg.matrix[k][u]
			}
			if flag[u] == true && weigth < dist[u] {
				dist[u] = weigth
				prev[u] = k
			}
		}

	}

	for i := 0; i < gg.vexnum; i++ {
		fmt.Printf("shortest %s->%s = %d\n", gg.vexs[start], gg.vexs[i], dist[i])
	}
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
