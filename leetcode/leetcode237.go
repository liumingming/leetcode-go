package leetcode

/*
* Definition for singly-linked list.这个原地删除节点的方法确实牛逼
* */
type ListNode struct {
    Val int
     Next *ListNode
 }

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
