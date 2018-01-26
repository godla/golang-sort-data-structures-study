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
	fBFS(&gg)
	fDFS(&gg)

	//listgg := list.New()

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

func prim() {

}
