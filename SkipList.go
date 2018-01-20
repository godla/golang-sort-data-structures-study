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

	fmt.Println(sl, sl.hnode.forward)
	for i := 0; i < 10; i++ {
		insert(sl, i)
	}

	fmt.Println(search(sl, 11))
	fmt.Println(search(sl, 1))
	fmt.Println(delete(sl, 1))
	fmt.Println(search(sl, 1))
}

func insert(sl *SkipList, kv int) bool {
	var p, q *node
	//var update []*node
	update := make([]*node, MAX_LEVEL, MAX_LEVEL)
	p = sl.hnode
	k := sl.lv

	for i := k - 1; i >= 0; i-- { //loop level
		for p != nil && p.forward[i] != nil && p.forward[i].key < kv {
			p = p.forward[i]
		}
		update[i] = p
	}

	if p != nil && p.key == kv {
		return false
	}

	k = randomLevel()

	//update hnode forword pointer
	if k > sl.lv {
		for i := sl.lv; i < k; i++ {
			update[i] = sl.hnode
		}
		sl.lv = k
	}

	q = createNode(k, kv)

	for i := 0; i < k; i++ {
		if update[i] != nil {
			q.forward[i] = update[i].forward[i]
			update[i].forward[i] = q
		} else {

		}
	}

	return true
}

func delete(sl *SkipList, kv int) bool {
	update := make([]*node, MAX_LEVEL, MAX_LEVEL)
	p := sl.hnode

	for i := sl.lv - 1; i >= 0; i-- {
		for p != nil && p.forward[i] != nil && p.forward[i].key < kv {
			p = p.forward[i]
		}
		update[i] = p
	}

	if p.forward != nil && p.forward[0].key != kv {

		return false
	}

	deleteN := p.forward[0]
	for i := sl.lv - 1; i >= 0; i-- {
		if update[i] != nil && update[i].forward[i] == deleteN {
			update[i].forward[i] = deleteN.forward[i]
			if sl.hnode.forward[i] == nil {
				sl.lv--
			}
		}
	}
	return true
}

func search(sl *SkipList, kv int) int {
	var q *node
	q = sl.hnode
	for i := sl.lv - 1; i >= 0; i-- {
		for q != nil && q.forward[i] != nil && q.forward[i].key <= kv {
			if q.forward[i].key == kv {
				return q.forward[i].key
			}
			q = q.forward[i]
		}
	}
	return -1
}
