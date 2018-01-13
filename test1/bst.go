package main

import (
	"fmt"
)

func main() {
	var data = []int{6, 5, 3, 1, 8, 7, 2, 4, 9}

	//
	var tree *bst
	fmt.Println(tree)
	tree = insert(tree, data[0])

	//init tree
	for _, v := range data {
		tree = insert(tree, v)
	}

	printTree(tree, 0, "")
	for i, str := range fstr {
		fmt.Println("L", i, str)
	}

	// for i := len(data) - 1; i >= 0; i-- {
	// 	fmt.Println("search", data[i])
	// 	search(tree, data[i])
	// }

	tree = tdelete(tree, 8)

	fstr = make([]string, 20, 20)
	printTree(tree, 0, "")
	for i, str := range fstr {
		fmt.Println("L", i, str)
	}
}

type bst struct {
	v int
	l *bst
	r *bst
}

func insert(tree *bst, v int) *bst {
	if tree == nil {
		tree = new(bst)
		tree.v = v
		return tree
	}

	if v > tree.v {
		tree.r = insert(tree.r, v)
	}

	if v < tree.v {
		tree.l = insert(tree.l, v)
	}

	if v == tree.v {
		fmt.Println("error")
	}
	return tree
}

func tdelete(tree *bst, v int) *bst {
	if tree == nil {
		return tree
	}

	if v < tree.v {
		tree.l = tdelete(tree.l, v)
	}
	if v > tree.v {
		tree.r = tdelete(tree.r, v)
	}

	if tree.v == v {
		tree = deleteR(tree)
	}

	return tree
}

func deleteR(tree *bst) *bst {
	if tree.l == nil && tree.r == nil {
		fmt.Println("tree.l == nil && tree.r == nil")
		return nil
	} else if tree.l != nil && tree.r == nil {
		return tree.l
	} else if tree.r != nil && tree.l == nil {
		return tree.r
	} else if tree.r != nil && tree.l != nil {
		lc := tree.l
		var s *bst
		for lc.r != nil {
			s = lc.r
		}
		if s == nil {
			lc.r = tree.r
		} else {
			s.r = tree.r
		}
		tree = tree.l
	}
	return tree
}

func search(tree *bst, v int) {
	if tree == nil {
		fmt.Println("tree == nil")
		return
	}

	if v == tree.v {
		fmt.Println("find", v)
	}

	if v > tree.v {
		search(tree.r, v)
	}

	if v < tree.v {
		search(tree.l, v)
	}
}

//for test view
var fstr = make([]string, 20, 20)

func printTree(tree *bst, i int, n string) {
	i++
	str := " "
	for n := i; n < 9; n++ {
		str += "-"
	}
	tmp := fmt.Sprintf("%s[%d]%s%s", str, tree.v, n, str)
	fstr[i] += tmp
	if tree.l != nil {
		printTree(tree.l, i, "L"+fmt.Sprintf("%d", tree.v))
	}
	if tree.r != nil {
		printTree(tree.r, i, "R"+fmt.Sprintf("%d", tree.v))
	}
}
