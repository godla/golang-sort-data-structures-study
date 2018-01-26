package main

import (
	"container/list"
	"fmt"
)

const MAX_SIZE int = 5

//为了看上去 好一些
const MAX_VALUE int = 9

func main() {
	fmt.Println("Prim")
	var gg Graph
	var vexs = []string{"B", "A", "C", "D", "E"}
	gg.vexnum = 5
	gg.vexs = vexs

	for i := 0; i < len(vexs); i++ {
		for j := 0; j < len(vexs); j++ {
			gg.matrix[i][j] = MAX_VALUE
		}
	}
	initGG(&gg)
	fmt.Println(gg.vexs)
	fBFS(&gg)
	fDFS(&gg)

	//listgg := list.New()
	prim(&gg, 0)
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

	gg.edgnum = 12 / 2
}

//深度遍历
func DFS(gg *Graph, visit *[]bool, i int) {

	fmt.Println(gg.vexs[i])
	for j := 0; j < gg.vexnum; j++ {
		if gg.matrix[i][j] != MAX_VALUE && !(*visit)[j] {
			(*visit)[j] = true
			DFS(gg, visit, j)
		}
	}
}

func fDFS(gg *Graph) {
	visit := make([]bool, 10, 10)
	fmt.Println(visit)
	visit[0] = true
	DFS(gg, &visit, 0)
}

//广度遍历
func fBFS(gg *Graph) {
	listq := list.New()
	visit := make([]bool, 10, 10)

	//first push
	visit[0] = true
	listq.PushBack(0)

	for listq.Len() > 0 {
		index := listq.Front()
		fmt.Println(gg.vexs[index.Value.(int)])
		for i := 0; i < gg.vexnum; i++ {
			if !visit[i] && gg.matrix[index.Value.(int)][i] != MAX_VALUE {
				visit[i] = true
				listq.PushBack(i)
			}
		}
		listq.Remove(index)
	}
}

func prim(gg *Graph, start int) {
	index := 0
	sum := 0
	prims := make([]string, 10, 10)
	var weights [5][2]int //[[0 0] [0 5] [0 3] [0 9] [0 9]]

	prims[index] = gg.vexs[start]
	index++

	//next vex
	for i := 0; i < gg.vexnum; i++ {
		weights[i][0] = start               //k
		weights[i][1] = gg.matrix[start][i] //v
	}

	//delete vex
	weights[start][1] = 0

	for i := 0; i < gg.vexnum; i++ {
		//fmt.Println(weights)
		if start == i {
			continue
		}

		min := MAX_VALUE
		next := 0
		for j := 0; j < gg.vexnum; j++ {
			if weights[j][1] != 0 && weights[j][1] < min {
				min = weights[j][1]
				next = j
			}
		}

		fmt.Println(gg.vexs[weights[next][0]], gg.vexs[next], "权重", weights[next][1])
		sum += weights[next][1]
		prims[index] = gg.vexs[next]
		index++

		//delete vex
		weights[next][1] = 0

		//update
		for j := 0; j < gg.vexnum; j++ {
			if weights[j][1] != 0 && gg.matrix[next][j] < weights[j][1] {
				weights[j][1] = gg.matrix[next][j]
				weights[j][0] = next
			}
		}
	}

	fmt.Println("sum:", sum)
	fmt.Println(prims)
}

func get_position(gg *Graph, ch string) int {
	for i := 0; i < gg.vexnum; i++ {
		if gg.vexs[i] == ch {
			return i
		}
	}
	return -1
}
