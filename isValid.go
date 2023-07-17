package main

import (
	"container/list"
	"fmt"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。

func main() {
	s := "]["
	result := isValid(s)
	fmt.Println(result)
}

func isValid(s string) bool {
	bracket := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := list.New()
	if len(s)%2 == 1 {
		return false
	}
	for _, v := range s {
		if string(v) == "(" || string(v) == "[" || string(v) == "{" {
			stack.PushBack(string(v))
		} else {
			if stack.Len() == 0 {
				return false
			}
			back := stack.Back().Value
			stack.Remove(stack.Back())
			if back != bracket[string(v)] {
				return false
			}
		}
	}
	isEmpty := stack.Len() == 0
	if isEmpty == true {
		return true
	}
	return false
}
