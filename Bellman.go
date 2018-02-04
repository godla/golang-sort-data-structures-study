package main

import (
	"fmt"
)

type Edge struct {
	u, v, weight int
}

var edge [10]Edge
var dist [10]int
var maxValue int

var source int
var nodeNum int
var edgeNum int

func init() {
	maxValue = 100
}

func initEdge() {
	fmt.Println("input nodeNum edgeNum source")
	fmt.Scanf("%d %d %d", &nodeNum, &edgeNum, &source)
	fmt.Println(nodeNum, edgeNum, source)

	for i := 0; i <= nodeNum; i++ {
		dist[i] = maxValue
	}
	dist[source] = 0

	for i := 0; i < edgeNum; i++ {
		fmt.Println("input edge.strat edge.end edge.weight")
		fmt.Scanf("%d %d %d", &edge[i].u, &edge[i].v, &edge[i].weight)
		if edge[i].u == source {
			dist[edge[i].v] = edge[i].weight
		}
	}
	fmt.Println(edge)
	fmt.Println(dist)
}

func Bellman() {
	for i := 0; i < nodeNum-1; i++ {
		for j := 0; j < edgeNum; j++ {
			//开始节点 权值 >
			if dist[edge[j].v] > dist[edge[j].u]+edge[j].weight {
				dist[edge[j].v] = dist[edge[j].u] + edge[j].weight
			}
			fmt.Println(dist)
		}
	}

}

func main() {
	//Bellman 差分约束系统 线性规划
	fmt.Println("Bellman")
	initEdge()
	Bellman()
}
