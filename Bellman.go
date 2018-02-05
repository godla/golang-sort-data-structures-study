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
	//nodeNum -1 自身节点需要去除
	for i := 0; i < nodeNum-1; i++ {
		//查找已知权重节点 相连接节点 并且更新权重
		for j := 0; j < edgeNum; j++ {
			if dist[edge[j].v] > dist[edge[j].u]+edge[j].weight {
				dist[edge[j].v] = dist[edge[j].u] + edge[j].weight
			}

		}
	}
	fmt.Println(dist)

	//不存在负环路时，都有 v.d < = u.d + w ( u , v )
	for i := 0; i < edgeNum; i++ {
		if dist[edge[i].v] > dist[edge[i].u]+edge[i].weight {
			//存在负环路时，一定存在某条边使得 v.d >u.d + w ( u , v )
			fmt.Println("Find 负环路")
			return
		}
	}
	//另一种方案 从start出发。不断维护每个点的最短距离，如果有负权环，则会进行无数次的维护，越来越小，所以如果循环次数大于了V - 1则有负权环。
}

func main() {
	//Bellman 差分约束系统 线性规划
	fmt.Println("Bellman")
	initEdge()
	Bellman()
}
