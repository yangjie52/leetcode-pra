package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	cur := root
	var arr []int
	for cur != nil {
		arr = append(arr)
		fmt.Println(cur.Val)
		cur = inorderTraversal(cur.Right)
	}
	return arr
}

func main() {

}
