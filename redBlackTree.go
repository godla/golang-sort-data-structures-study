package main

import (
	"fmt"
)

func main() {
	fmt.Println("red-black-tree")

	var n node
	(&n).rotateR()
	fmt.Println(n)
	fmt.Println(n.r)
}

type node struct {
	d int
	c bool //0 = red 1=black
	l *node
	r *node
	p *node
}

func insert() {

}

func delte() {

}

func search() {

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

// rotate cur node
func (n *node) rotateR() {
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

	if gp != nil {
		if gp.l == p {
			gp.l = n
		} else {
			gp.r = n
		}
	}
}

func (n *node) rotateL() {
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
