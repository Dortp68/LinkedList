package List

import "fmt"

type Node struct{
	data_ int
	pnext *Node
}
type List struct{
	head *Node
	size_ int
}

func NewList() *List{
	return &List{nil, 0}
}

func (l *List) GetSize() int{
	return l.size_
}

func (l *List) Append(data_ int){
	switch {
	case l.head == nil:
		tmp := new(Node)
		tmp.data_ = data_
		l.head = tmp
		l.size_++
	default:
		cur := l.head
		for j:=0; j!=l.size_-1; j++{
			cur = cur.pnext
		}
		tmp := new(Node)
		tmp.data_ = data_
		cur.pnext = tmp
		l.size_++
	}
}

func (l *List) Add(data_ int, id int) error{
	switch {
	case id<0 || id>l.size_:
		return fmt.Errorf("out of range")
	case id == 0:
		tmp := new(Node)
		tmp.data_ = data_
		tmp.pnext = l.head
		l.head = tmp
		l.size_++
		return nil
	default:
		cur := l.head
		for j:=0; j!=id-1; j++{
			cur = cur.pnext
		}
		tmp := new(Node)
		tmp.data_ = data_
		tmp.pnext = cur.pnext
		cur.pnext = tmp
		l.size_++
		return nil
	}
}

func (l *List) Remove(id int) error{
	switch {
	case id<0 || id>=l.size_:
		return fmt.Errorf("out of range")
	case id==0:
		l.head = l.head.pnext
		l.size_--
		return nil
	default:
		cur := l.head
		for j:=0; j!=id-1; j++{
			cur = cur.pnext
		}
		tmp:=cur.pnext
		cur.pnext = tmp.pnext
		l.size_--
		return nil
	}
}

func (l *List) PrintAll(){

	for cur := l.head; cur!=nil; cur=cur.pnext{
		fmt.Print(cur.data_, " ")
	}
	fmt.Print("\n")
}