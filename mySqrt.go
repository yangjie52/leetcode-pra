package main

import "fmt"

func mySqrt(x int) int {
	sqrt := 0
	for i := 0; i <= x; i++ {
		if i*i <= x && (i+1)*(i+1) > x {
			sqrt = i
			break
		}
	}
	return sqrt
}

func main() {
	x := 1
	result := mySqrt(x)
	fmt.Println(result)
}
