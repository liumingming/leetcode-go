package gotype

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
func ReverseNoHeadNode(node *ListNode) *ListNode{
	if node == nil || node.Next == nil {
		return node
	}

	var prev *ListNode
	var curr *ListNode

	curr = node

	for curr.Next != nil  {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return curr
}

/*
	1.1.2 从尾到头输出链表
*/
func ReversePrint(node *ListNode){
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
