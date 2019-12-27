package gotype

import (
	"sync"
)

type SliceStack struct {
	arr []interface{}
	stackSize int
	sync.RWMutex
}

func NewSliceStack() *SliceStack  {
	return &SliceStack{arr:make([]interface{}, 0)}
}

func (p *SliceStack)IsEmpty() bool {
	return p.stackSize == 0
}

func (p *SliceStack)Size() int {
	return p.stackSize
}

func (p *SliceStack)Top() interface{} {
	if  p.IsEmpty() {
		return nil
	}
	return p.arr[p.stackSize-1]
}

func (p *SliceStack)List() []interface{}  {
	return p.arr
}

func (p *SliceStack)Pop() interface{}  {
	p.Lock()
	defer p.Unlock()

	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}

	return nil
}

func (p *SliceStack)Push(t interface{})  {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}


type LinkedStack struct {
	head *ListNode
	sync.RWMutex
}

func NewLinkedStack()  *LinkedStack  {
	return &LinkedStack{head:&ListNode{}}
}

func (p *LinkedStack)IsEmpty() bool {
	return p.head.Next == nil
}

func (p *LinkedStack)Size() int {
	size := 0
	node := p.head.Next

	for node != nil {
		node = node.Next
		size++
	}

	return size
}

func (p *LinkedStack)Push(t interface{})  {
	p.Lock()
	defer p.Unlock()
	node := &ListNode{Data:t, Next:p.head.Next}
	p.head.Next = node
}


func (p *LinkedStack)Pop() interface{}  {
	p.Lock()
	defer p.Unlock()
	temp  := p.head.Next

	if temp != nil {
		p.head.Next = temp.Next
		return temp.Data
	}
	return nil
}

func (p *LinkedStack)Top() interface{}  {
	if p.head.Next != nil {
		return p.head.Next.Data
	}
	return nil
}



