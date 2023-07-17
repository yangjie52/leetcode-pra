package main

import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {
	var sum int
	var str string
	str = strconv.FormatUint(uint64(num), 2)
	//fmt.Println(str)
	for _, v := range str {
		fmt.Println(v)
		if string(v) == "1" {
			sum++
		}
	}
	return sum
}

func main() {
	var n uint32

	n = 00000000000000000000000010000000
	result := hammingWeight(n)
	fmt.Println(result)
}
