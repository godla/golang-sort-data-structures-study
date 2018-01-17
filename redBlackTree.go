package main

import (
	"fmt"
)

func main() {
	fmt.Println("red-black-tree")
	var data = []int{6, 5, 3, 1, 8, 7, 2, 4, 9, 0, 3}
	//var data = []int{6, 5, 3, 1}

	tree := NewTree()
	for _, v := range data {
		tree.insertB(v)
	}
	//fmt.Println(tree.size)
	tree.tdelete(tree.root, 2)
	printT(tree.root)
}

type node struct {
	l, r, p *node
	v       int
	c       bool //0 = red 1=black
}

type Tree struct {
	root *node
	last *node
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

func (n *node) getGp() *node {
	if n.p == nil {
		return nil
	}
	if n.p.p == nil {
		return nil
	}
	return n.p.p
}

func (n *node) getUn() *node {
	if n.getGp() == nil {
		return nil
	}
	if n.p == n.getGp().l {
		return n.getGp().r
	} else {
		return n.getGp().l
	}
}

// rotate cur node
func (tree *Tree) rotateR(n *node) {
	gp := n.getGp()
	p := n.p
	r := n.r

	n.r = p
	n.p = gp

	if p != nil {
		p.p = n
		p.l = r
		if r != nil {
			r.p = p
		}
	}
	if tree.root == p {
		tree.root = n
	}

	if gp != nil {
		if gp.l == p {
			gp.l = n
		} else {
			gp.r = n
		}
	}
}

func (tree *Tree) rotateL(n *node) {
	gp := n.getGp()
	p := n.p
	l := n.l

	n.l = p
	n.p = gp

	if p != nil {
		p.p = n
		p.r = l
		if l != nil {
			l.p = p
		}
	}

	if tree.root == p {
		tree.root = n
	}

	if gp != nil {
		if gp.l == p {
			gp.l = n
		} else {
			gp.r = n
		}
	}

}

func getMin() {

}

func getMax() {

}

func (tree *Tree) tdelete(n *node, v int) {
	if tree == nil {
		return
	}
	if v < n.v {
		tree.tdelete(n.l, v)
	}
	if v > n.v {
		tree.tdelete(n.r, v)
	}

	if n.v == v {
		if n.l != nil && n.r != nil {
			mn := tree.findMax(n)
			//转换为 mn.r 必定为 nil 依据为2叉搜寻树特性 小的放左边 大的放右边
			n.v = mn.v //swap data
			if tree.deleteR(mn) {
				tree.size--
			}
		} else {
			if tree.deleteR(n) {
				tree.size--
			}
		}

	}
}

//n.l != nil && n.r != nil 前置条件
func (Tree *Tree) findMax(n *node) *node {
	//选择n节点的左子树中最大节点
	maxn := n.l
	for maxn.r != nil {
		maxn = maxn.r
	}
	return maxn
}

//     n			   n        = delete n
//   n  nil   ==   nil   n
//nil nil             nil nil
func (tree *Tree) deleteR(n *node) bool {
	red := false
	black := true

	//case1 root
	if n.l == nil && n.r == nil && n.p == nil {
		n = nil
		tree.root = n
		return true
	}

	var child *node
	if n.l != nil {
		child = n.l
	} else {
		child = n.r
	}

	//case2 root
	if n.p == nil {
		child.p = nil
		tree.root = child
		child.c = true
		return true
	}

	if n.p.l == n {
		n.p.l = child
	} else {
		n.p.r = child
	}

	child.p = n.p

	if n.c == black {   
		if child != nil && child.c == red {
			//          p
			//       b 			= n
			//    r      nil
			// nil nil
			//fcase 1
			child.c = true
		} else {
 
			tree.fix(child)
		}
	}

	// 由于 红色节点采用直接删除 所以可能产生
	//       b
	//     b    b
	//   b    b   b
	//if red is ok

	//free n
	return true
}

//get delete node child
//will retrun nil
func getDc(n *node) *node {
	if n.l != nil {
		return n.l
	} else {
		return n.r
	}
}

//兄弟
func (n *node) br() *node {
	if n.p.l == n {
		return n.r
	} else {
		return n.l
	}
}

func cleanN(n *node){

}

func sibling(n *node){
	if n.p.l == n {
		return n.r
	} else {
		return n.l
	}
}

//n = delete node n=black
//更具红黑树 特性 推论出 如下结构 
//   n=b         
//nil   nil
//在这种情况下，删除n 将使得 整个tree 不平衡 少了一个黑节点嘛 
//所以思想是 从隔壁挪一个黑色过来 或者 重新染色
func (tree *Tree) fix2(n *node) {

	red := false
	black := true


	//余下情况 根据此图考虑  
	//因为n=black 情况下必然存在兄弟br 
	//那为什么有brl brr 如果 br 等于red 那brl brr = black
	//所以下面是总图 并非每种情况图 
	//          p
	//     n=b       br
	// nil  nil brl      brr
	//        nil nil nil nil
	//由于之前用br 发现描述有点难受 查了下资料发现都用 sibling br 我们改成s 
	//          p
	//    n=b       s
	// nil  nil  sl      sr
	//        nil nil nil nil
	//看上去很眼熟 是不是 没错大部分教程都是这个图
	
	//case1
	if n.p == nil{
		tree.root = n.l
		cleanN(n)
		return
	}

	//case2
	//          b
	//     b         r
	// nil  nil  b       b
	//        nil nil nil nil
	if  n.br().c == red && n.c == black && n.p.c == black {
		n.p.c = red
		n.br().c = black
		if n == n.p.l {
			tree.rotateL(n.br())
		} else {
			tree.rotateR(n.br())
		}
	//            b
	//       r         b
	//   b=n       b
	//nil nil nil nil  
	//这种情况下删除 n 并不能保证平衡 所以要继续执行case
	}

	//case 3 全黑
	//         b
	//    b            b
	// nil nil   b=nil    b=nil
	if n.p.c == black && n.br().c == black && n.br().l==nil && n.br().r == nil{
		n.br().c = red
		return
	}

	//case3-2
	//         b
	//    b              b
	// nil nil       r       r
	//           nil  nil  nil nil
	//fcase 4
	if n.p.c = red && n.br().c = black && n.br().l == black && n.br().r ==black{
		n.p.c = black
		n.br().c = red
	}


	//fcase 5
	if 
}

//修复红黑树平衡 n.c = black n.p.c = black
func (tree *Tree) fix(n *node) {
	// if n.p == nil {
	// 	n.c = true
	// 	tree.root = n
	// 	return
	// }
	red := false
	black := true
	//br = red
	if n.br().c == red {
		n.p.c = red
		n.br().c = black
		if n == n.p.l {
			tree.rotateL(n.br())
		} else {
			tree.rotateR(n.br())
		}
	}
	if n.p.c == black && n.br().c == black && n.br().l.c == black && n.br().r.c == black {
		n.br().c = red
		tree.fix(n.p) //????????????
	} else if n.p.c == red && n.br().c == black && n.br().l.c == black && n.br().r.c == black {
		n.br().c = red
		n.p.c = black
	} else {
		if n.br().c == black {
			if n == n.p.l && n.br().l.c == red && n.br().r.c == black {
				n.br().c = red
				n.br().l.c = black
				tree.rotateR(n.br().l)
			} else if n == n.p.r && n.br().l.c == black && n.br().r.c == red {
				n.br().c = red
				n.br().r.c = black
				tree.rotateL(n.br().r)
			}
		}
		n.br().c = n.p.c
		n.p.c = black
		if n == n.p.l {
			n.br().r.c = black
			tree.rotateL(n.br())
		} else {
			n.br().l.c = black
			tree.rotateR(n.br())
		}
	}

}

func (tree *Tree) insertB(v int) {
	if tree.root == nil {
		tree.root = new(node)
		tree.root.v = v
		tree.root.c = true
		tree.size++
	}

	if v < tree.root.v {
		if tree.Insert(&tree.root.l, v, tree.root) {
			printT(tree.root)
			tree.size++
			tree.inserCase(tree.last)
		}
	}

	if v > tree.root.v {
		if tree.Insert(&tree.root.r, v, tree.root) {
			printT(tree.root)
			tree.size++
			tree.inserCase(tree.last)
		}
	}
	printT(tree.root)
}

func (tree *Tree) Insert(n **node, v int, fa *node) bool {

	pn := (*n)
	if (*n) == nil {
		(*n) = new(node)
		(*n).v = v
		(*n).p = fa
		tree.last = (*n)
		return true
	}

	if v > pn.v {
		tree.Insert(&(pn.r), v, *n)
	}

	if v < pn.v {
		tree.Insert(&(pn.l), v, *n)
	}

	if v == pn.v {
		return false
	}
	return true
}

func (tree *Tree) inserCase(n *node) {
	if n.p == nil {
		n.c = true
		tree.root = n
		return
	}

	if n.p.c == false {
		if n.getUn() != nil && n.getUn().c == false {
			n.p.c = true
			n.getUn().c = true
			n.getGp().c = false
			tree.inserCase(n.getGp()) //if root node
		} else {
			//nil or black
			if n == n.p.r && n.getGp() != nil && n.p == n.getGp().l {
				tree.rotateL(n)
				tree.rotateR(n)
				n.c = true
				n.l.c = false
				n.r.c = false
			}
			if n == n.p.l && n.getGp() != nil && n.p == n.getGp().r {
				tree.rotateR(n)
				tree.rotateL(n)
				n.c = true
				n.l.c = false
				n.r.c = false
			}
			if n == n.p.l && n.getGp() != nil && n.p == n.getGp().l {

				n.p.c = true
				if n.getGp() != nil {
					n.getGp().c = false
				}
				tree.rotateR(n.p)
			}
			if n == n.p.r && n.getGp() != nil && n.p == n.getGp().r {
				n.p.c = true
				if n.getGp() != nil {
					n.getGp().c = false
				}
				tree.rotateL(n.p)
			}
		}

	}
}

//test print ----------------------------------------------
var fstr = make([]string, 20, 20)

func printT(tree *node) {
	fmt.Println(tree)
	fstr = make([]string, 8, 8)
	printTree(tree, 0, "")
	for i, str := range fstr {
		fmt.Println("L", i, str)
	}
}

func printTree(tree *node, i int, n string) {
	fmt.Println("-----", tree)
	i++
	str := " "
	for n := i; n < 9; n++ {
		str += "-"
	}
	var tmp string
	if tree.c == true {
		tmp = fmt.Sprintf("%s\033[40;37m [%d] \033[0m%s%s", str, tree.v, n, str)
	} else {
		tmp = fmt.Sprintf("%s\033[41;37m [%d] \033[0m%s%s", str, tree.v, n, str)
	}

	fstr[i] += tmp
	if tree.l != nil {
		printTree(tree.l, i, "L"+fmt.Sprintf("%d", tree.v))
	}
	if tree.r != nil {
		printTree(tree.r, i, "R"+fmt.Sprintf("%d", tree.v))
	}
}
