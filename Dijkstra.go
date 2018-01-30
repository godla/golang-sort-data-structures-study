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
	var ss [MAX_SIZE]bool  //最短路劲节点集合

	//init
	dist = gg.matrix[start]
	ss[start] = true //find start to start
	dist[start] = 0  //start to start length

	for i := 0; i < gg.vexnum; i++ {
		k := 0
		min := MAX_VALUE
		fmt.Println("-----------")
		fmt.Println(dist, ss)
		//find next 贪心
		for j := 0; j < len(dist); j++ {
			if ss[j] == false && dist[j] != MAX_VALUE && dist[j] < min {
				min = dist[j]
				k = j
			}
		}

		//set find
		ss[k] = true

		//update dist length
		for u := 0; u < gg.vexnum; u++ {
			if gg.matrix[k][u] != MAX_VALUE && ss[u] == false {
				weight := min + gg.matrix[k][u]
				if weight < dist[u] {
					dist[u] = weight
				}
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
	gg.matrix[1][3] = 99
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
