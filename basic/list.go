package basic

import "fmt"

/*单链表*/
type SingleList struct {
	head *ListNode
}

type ListNode struct {
	data int
	next *ListNode
}

//新建单链表
func NewSingleList() *SingleList{
	head := &ListNode{}
	return &SingleList{head:head}
}

//在头节点处增加节点
func (s *SingleList) AddNode(data int) {
	node := &ListNode{data: data, next:s.head.next}
	s.head.next = node
}

func (s *SingleList) Print() {
	curr := s.head.next
	for  curr.next != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println(curr.data)  //尾节点数据
}

//单链表反转
func (s *SingleList) Reverse(){
	curr := s.head.next
	var prev *ListNode
	for  curr.next != nil {
		curr.next, prev, curr = prev, curr, curr.next
	}
	curr.next = prev
	s.head = curr
}

func (s *SingleList) Delete(node *ListNode){
	curr := s.head.next
	for  curr.next != nil {
		if curr.data == node.data {
			curr.next = curr.next.next
		}
	}
}