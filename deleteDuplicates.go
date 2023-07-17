package main

import "fmt"

// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //判断链表为空或者链表只有一个节点
		return head
	}
	cur := head //将cur初始化为链表的头节点head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next

		} else {
			cur = cur.Next
		}
		//fmt.Println(cur.Val)
	}
	return head
}

func main() {
	list1 := ListNode{}
	node11 := ListNode{Val: 1}
	node12 := ListNode{Val: 1}
	node13 := ListNode{Val: 4}
	node11.Next = &node12
	node12.Next = &node13
	list1.Next = &node11
	l1 := &list1
	result := deleteDuplicates(l1)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
