package main

import (
	"fmt"
)
const entry_size =10

func InitIndex() *Index {
	n := Node{A:1.0, b:0.0}
	ind := &Index{n}
	return ind
}

type Index struct {
	root Node
}

func (ind Index) Lookup(k uint32) (*Entry, bool) {
	node_ := &ind.root
	e := node_.forward(k)
	for e.getType() == 1 { // 2: Null, 0: Data, 1:Node
		node_ = e.d_.(NodeType).p_
		e = node_.forward(k)
	}
	if e.getType() == 0 {
		ek := e.d_.(DataType).k_
		if ek == k {
			return e, true
		}		
	}
	return e, false
}

func (ind Index) Insert(k uint32, p T) {
	e, _ := ind.Lookup(k)
	if e.getType() == 2 {
		e.t_ = 0
		d := DataType{k, p}
		e.d_ = d
	} else {
		ek := e.d_.(DataType).k_
		n := trainNode(k, ek)
		e.setType(1)
		e.d_ = NodeType{n}
	}	
}

func trainNode(k1, k2 uint32) *Node {
	n := Node{A:1.0, b:0.0}
	return &n
}

type Node struct {
	A float32
	b float32
	e [entry_size]Entry
}

func (n Node) forward(k uint32) *Entry {
	idx := int(n.A * float32(k) + n.b)

	if len(n.e) - 1 <= idx {
		return &n.e[len(n.e) - 1]
	} else if idx <= 0 {
		return &n.e[0]
	}	
	return &n.e[idx]
}


type Entry struct {
	t_ uint8
	d_ EntryType
}

func (e Entry) getType() uint8 {
	return e.t_
}

func (e Entry) setType(t uint8) {
	e.t_ = t
}

type EntryType interface {

}

type DataType struct {
	k_ uint32
	v_ T
}


type NodeType struct {
	p_ *Node 
}




func main() {
	fmt.Println("test")
	ind := InitIndex()
	fmt.Printf("%d\n", 1)
}


type T interface {
}
