package main

import (
	"container/list"
	"fmt"
)
const entry_size = 1000
const beta = 2
const alpha = 0.1

func InitIndex() *Index {
	n := Node{A:1.0, b:0.0}
	ind := &Index{n}
	return ind
}

type Index struct {
	root Node
}

func (ind *Index) Lookup(k uint32) (*Entry, bool) {
	node_ := &ind.root
	e := node_.forward(k)
	for e.getType() == 3 { // 0: Null, 2: Data, 3:Node
		node_ = e.d_.(NodeType).p_
		e = node_.forward(k)
	}
	if e.getType() == 2 {
		ek := e.d_.(DataType).k_
		if ek == k {
			return e, true
		}		
	}
	return e, false
}

func (ind *Index) lookupPath(k uint32) (*Entry, bool, *list.List) {
	path := list.New()
	node_ := &ind.root
	path.PushBack(node_)
	e := node_.forward(k)
	for e.getType() == 3 { // 0: Null, 2: Data, 3:Node
		node_ = e.d_.(NodeType).p_
		e = node_.forward(k)
		path.PushBack(node_)
	}
	if e.getType() == 2 {
		ek := e.d_.(DataType).k_
		if ek == k {
			return e, true, path
		}		
	}
	return e, false, path
}


func (ind *Index) Insert(k uint32, p T) bool {
	e, chk, path := ind.lookupPath(k)
	if chk {
		fmt.Println("Same key exists")
		return false
	}

	if e.getType() == 0 {
		e.setType(2)
		d := DataType{k, p}
		e.d_ = d
	} else {
		ek := e.d_.(DataType).k_
		
		A_, b_:= trainNode([]uint32{k, ek}, entry_size)
		n := &Node{A:A_, b:b_}
		e.setType(3)
		e.d_ = NodeType{n}
	}	

	chk = true
	for p := path.Back() ; p != nil ; p = p.Prev() {
		if p.Prev() != nil {
			e = p.Prev().Value.(*Node).forward(k)
			ind.adjust(p.Value.(*Node), e, chk)
		} else {
			ind.adjust(p.Value.(*Node), nil, chk)
		}		
		chk = false
	}

	return true
}

func (ind *Index) adjust(n *Node, pe *Entry, chk bool) {
	n.element_num += 1
	if chk {
		n.conflict_num += 1
	}
	
	if n.element_num >= beta * n.build_num && 
		float32(n.conflict_num) / float32(n.element_num - n.build_num) >= alpha {
			
		pe.d_ = buildPartialTree(nil)
	}
}

func buildPartialTree(ks []uint32) *Node {
	L := entry_size
	if L < 2 * len(ks) {
		L = 2 * len(ks)
	}
	A_, b_ := trainNode(ks, uint32(L))
	n := &Node{A:A_, b:b_}
	for i := 0 ; i < L ; i += 1 {
		subKs := []uint32{}
		if len(subKs) == 1 {
			n.e[i].setType(2)
			n.e[i].d_ = DataType{subKs[0], i}
		} else if len(subKs) > 1 {
			n.e[i].setType(3)
			n.e[i].d_ = NodeType{buildPartialTree(subKs)}
		}
	}
	n.element_num = uint32(len(ks))
	n.element_num = uint32(len(ks))
	n.conflict_num = 0
	return n
}

func trainNode(ks []uint32, L uint32) (float32, float32) {
	//FMCD
	i := 0
	T := 1
	N := len(ks)
	Ut :=  float32(ks[N-1-T] - ks[T]) / float32(L - 2)
	for i <= N - 1 - T {
		for i + T < N && float32(ks[i+T] - ks[i]) >= Ut {
			i += 1
		}

		if i + T >= N {
			break
		}

		T += 1
		Ut = float32(ks[N-1-T] - ks[T]) / float32(L - 2)
	}

	A_ := 1 / Ut
	b_ := (entry_size - (A_ * float32(ks[N-1-T] + ks[T])))/ 2
	return A_, b_
}

type Node struct {
	A float32
	b float32
	element_num uint32
	conflict_num uint32
	build_num uint32
	e [entry_size]Entry
}

func (n *Node) forward(k uint32) *Entry {
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

func (e *Entry) getType() uint8 {
	return e.t_
}

func (e *Entry) setType(t uint8) {
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
	_, chk := ind.Lookup(2)
	fmt.Println(chk)
	ind.Insert(2, 3)
	_, chk = ind.Lookup(2)
	fmt.Println(chk)
	ind.Insert(11, 3)
	ind.Insert(12, 3)
	
}


type T interface {
}
