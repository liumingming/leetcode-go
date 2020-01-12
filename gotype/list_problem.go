package gotype

import "fmt"

/*
	1.1给定一个带头结点的单链表，请将其逆序，即如果单链表原来为head->1->2->3,则逆序后变为head 3->2->1
 */
func Reverse(node *ListNode){
	if node == nil || node.Next == nil {
		return
	}

	var pre *ListNode
	var curr *ListNode
	var next *ListNode
	next = node.Next

	for next != nil  {
		curr = next.Next
		next.Next = pre
		pre = next
		next = curr
	}
	node.Next = pre
}


/*
	1.1.1 对不带头节点的单链表，请将其逆序，即如果单链表原来为1->2->3,则逆序后变为3->2->1
*/
func ReverseNoHeadNode(head *ListNode) *ListNode{

	var prev *ListNode
	var curr *ListNode = head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

/*
	1.1.2 从尾到头输出链表
*/
func ReversePrint(node *ListNode){
	if node == nil {
		return
	}

	ReversePrint(node.Next)

	fmt.Print(node.Data, " ")
}

/*
1.2 如何从无序链表中移除重复项
给定一个没有排序的链表，去掉其中重复项，并保留原顺序，例如链表1-》3-》1-》5-》7去掉重复项后变为1-》3=》5=》7
 */
func RemoveDup(head *ListNode)  {

}

/*
1.3 如何计算两个单链表所代表的数之和
 */
func Add(h1 *ListNode, h2  *ListNode) *ListNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}

	if h2 == nil || h2.Next == nil {
		return h1
	}

	c := 0
	sum := 0
	p1 := h1.Next
	p2 := h2.Next
	result := &ListNode{}
	p := result
	for p1 != nil && p2 != nil {
		p.Next = &ListNode{}
		sum = p1.Data.(int) + p2.Data.(int)
		p.Next.Data = sum % 10
		c = sum /10
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1 == nil {
		for p2 != nil {
			p.Next = &ListNode{}
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p2 = p2.Next
		}
	}

	if p2 == nil {
		for p1 != nil {
			p.Next = &ListNode{}
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p1 = p1.Next
		}
	}

	if c == 1 {

		p.Next = &ListNode{}
		p.Next.Data = 1
	}

	return result
}

/*

 */



