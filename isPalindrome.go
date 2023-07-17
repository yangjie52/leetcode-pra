package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a int
	a = 11
	//c := strconv.Itoa(a)3
	//result := isPalindrome(a)
	//isPalindrome2(a)
	result := isPalindrome4(a)
	fmt.Println(result)
}
func isPalindrome4(x int) bool {
	stack := list.New()
	stringx := strconv.Itoa(x)
	length := len(stringx)
	var slice int
	var invert string
	if x < 0 {
		return false
	}
	//入栈
	for i := 0; i < length/2; i++ {
		stack.PushBack(string(stringx[i]))
		//fmt.Println(stringx[i])
	}
	//出栈
	for stack.Len() != 0 {
		pop := stack.Back().Value.(string)
		invert += pop
		stack.Remove(stack.Back())
	}
	if length%2 == 1 {
		slice = length - length/2
	} else {
		slice = length / 2
	}
	fmt.Println(slice)
	if invert == stringx[slice:] {
		fmt.Println(stringx[slice:])
		return true
	} else {
		fmt.Println(stringx[slice:])
		return false
	}

}
func isPalindrome3(x int) bool {
	stack := list.New()
	stringX := strconv.Itoa(x)
	var invert string
	for _, v := range stringX {
		stack.PushBack(string(v))
	}
	for stack.Len() != 0 {
		pop := stack.Back().Value.(string)
		invert += pop
		stack.Remove(stack.Back())
	}
	if invert == stringX {
		return true
	} else {
		return false
	}
}
func isPalindrome(x int) bool {
	stringX := strconv.Itoa(x)
	t := x
	xlen := len(stringX)
	var compareX int
	if t < 0 {
		return false
	}
	for i := xlen; i > 0; i-- {
		x1 := math.Pow(10, float64(i-1)) //10的次方
		result := int(float64(t) / x1)
		t = t - (result * int(x1))
		x2 := math.Pow(10, float64(xlen-i))
		compareX += result * int(x2)

	}
	if compareX == x {
		return true
	} else {
		return false
	}
}

func isPalindrome1(x int) bool {
	stringX := strconv.Itoa(x)
	runeX := []rune(stringX)
	var ComX string
	for i := len(runeX) - 1; i >= 0; i-- {
		ComX += string(runeX[i])
	}
	//fmt.Println(ComX)
	if ComX == stringX {
		return true
	} else {
		return false
	}
}

func isPalindrome2(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}
	y := x
	comX := 0
	for x > comX {
		comX = comX*10 + y%10
		y /= 10

	}
	fmt.Println(comX)
	if comX == x {
		return true
	} else {
		return false
	}
}
