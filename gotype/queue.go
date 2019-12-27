package gotype

import "sync"

type SliceQueue struct {
	arr []interface{}
	sync.RWMutex
}

func NewSliceQueue()  *SliceQueue {
	return &SliceQueue{arr:make([]interface{}, 0)}
}

func (p *SliceQueue)IsEmpty() bool {
	return p.Size() == 0
}

func (p *SliceQueue)Size() int  {
	return len(p.arr)
}

func (p *SliceQueue)GetFront() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.arr[0]
}

func (p *SliceQueue)GetBack() interface{}{
	if p.IsEmpty() {
		return nil
	}
	return p.arr[p.Size()-1]
}





