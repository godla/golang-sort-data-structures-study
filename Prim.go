package main

import (
	"fmt"
)

var MAX_SIZE = 1000

func main() {
	fmt.Println("Prim")
}

type Graph struct {
	vexs   [MAX_SIZE]string        //定点集合
	vexnum int                     //定点数量
	edgnum int                     //边数量
	matri  [MAX_SIZE][MAX_SIZE]int //邻接矩阵
}

func init() {

}

func prim() {

}
