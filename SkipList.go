package main

import (
	"fmt"
	"math/rand"
)

var MAX_LEVEL int

type node struct {
	key     int
	forward []*node
}

type SkipList struct {
	lv    int //level
	hnode *node
}

func randomLevel() int {
	return rand.Intn(10)
}

func init() {
	MAX_LEVEL = 10
}

func createNode(lv int, kv int) *node {
	return &node{
		key:     kv,
		forward: make([]*node, lv, lv),
	}
}

func createSkipList() *SkipList {
	sl := new(SkipList)
	sl.lv = 0
	sl.hnode = createNode(MAX_LEVEL, 0)

	//象征性得初始化下
	// for i := 0; i < MAX_LEVEL; i++ {
	// 	sl.hnode.forward[i] = nil
	// }
	return sl
}

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("skip list", data)

	sl := createSkipList()
	fmt.Println(sl.hnode.forward)
}

func insert(sl *SkipList, kv int) bool {
	var p, q *node
	//var update []*node
	update := make([]*node, MAX_LEVEL, MAX_LEVEL)
	p = sl.hnode
	k := sl.lv

	for i := k - 1; i >= 0; i-- {
		q = p.forward[i]
		for q.key < kv {
			p = q
		}
		update[i] = p
	}

	//插入数据相同 直接返回false
	if q != nil && q.key == kv {
		return false
	}

	k = randomLevel()
	if k > sl.lv {
		for i := sl.lv; i < k; i++ {
			update[i] = sl.hnode
		}
		sl.lv = k
	}

	q = createNode(k, kv)
	for i := 0; i < k; i++ {
		q.forward[i] = update[i].forward[i]
		update[i].forward[i] = q
	}

	return true
}

func delete() {

}

func search(sl *SkipList, kv int) int {
	var p, q *node
	p = sl.hnode
	k := sl.lv
	for i := k - 1; i >= 0; i-- {
		q = p.forward[i]
		for q.key <= kv {
			if q.key == kv {
				return q.key
			}
			p = q
		}
	}
	return -1
}
