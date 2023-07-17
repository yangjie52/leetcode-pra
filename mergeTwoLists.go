package main

import "fmt"

//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{} //创建链表的头节点
	cur := head

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
			cur = cur.Next
		}
		if list1 != nil {
			cur.Next = list1
			break
		} else {
			cur.Next = list2
			break
		}
	}
	return head.Next
}

func main() {
	//list1 := ListNode{}
	//list2 := ListNode{}
	list1 := ListNode{}
	node11 := ListNode{Val: 1}
	node12 := ListNode{Val: 2}
	node13 := ListNode{Val: 4}
	node11.Next = &node12
	node12.Next = &node13
	list2 := ListNode{}
	node21 := ListNode{Val: 1}
	node22 := ListNode{Val: 2}
	node23 := ListNode{Val: 4}
	node21.Next = &node22
	node22.Next = &node23
	l1 := &list1
	l2 := &list2
	result := mergeTwoLists(l1, l2)
	fmt.Println(result)
}
