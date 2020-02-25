//package main
//
//
//type ListNode struct {
//	data int
//	Next *ListNode
//}
//func hasCycle(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//	return false
//	}
//
//	var fast *ListNode = head
//	var slow *ListNode = head
//
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//		if slow == fast {
//			return true
//		}
//	}
//
//	return false
//}
//
//func maxSubArray(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	var sum []int = make([]int, len(nums))
//	sum[0] = nums[0]
//	var max int = nums[0]
//
//	for i:=1; i<len(nums); i++ {
//		if sum[i-1] <=  0 {
//			sum[i] = nums[i]
//		} else {
//			sum[i] = nums[i] + sum[i-1]
//
//		}
//		max = int(math.Max(float64(max), float64(sum[i])))
//	}
//
//	return max
//}
//
//
//微信盆友圈
//
//点赞服务
//
//1 点赞是个写接口  10w：
//2 读点赞人
//
//容量规划：
//	qps多少 10亿日活， 没人每天点赞10次 美妙10万次  1000亿次  。百万qps
//	服务器  单机 5000 qps
//	存储
// 	代理层
//
//接入层
//    收到一个用户点赞
//
//业务层()
//	异步写入
//
//存储层
//	Redis集群
//	消息队列多盘
//
//监控
//
//报警


