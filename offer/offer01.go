package offer

import "fmt"

/*剑指offer面试题
面试题5：从尾到头打印链表
*/
type ListNode struct {
	value int
	next *ListNode
}

func PrintListReversely(head *ListNode)  {
	if head == nil {
		return
	}

	PrintListReversely(head.next)

	fmt.Println(head.value)
}
