package main

import (
	"fmt"
)

//示例 1：

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func main() {
	var strs []string = []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
	//fmt.Println(reflect.DeepEqual("ab", "ab"))
}

func longestCommonPrefix(strs []string) string {
	var shortest string
	var length int

	length = 201
	for i, v := range strs {
		if length >= len(v) {
			length = len(v)
			shortest = strs[i]
		}
	}
	fmt.Printf("最短的：%d 是：%s\n", length, shortest)
	for i := 0; i < length; i++ {
		for _, v1 := range strs {
			if i > length || v1[i] != shortest[i] {
				return shortest[:i]
			}

		}
	}
	return shortest
}
