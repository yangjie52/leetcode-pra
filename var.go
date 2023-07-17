package main

import "fmt"

func main() {
	var n int = 6
	for i := 0; i <= n; i++ {
		v1 := fibonacci(i)
		fmt.Printf("%d  ", v1)
	}

}

func factorial(n int) int {
	if n > 0 {
		result := n * factorial(n-1)
		return result
	} else {
		return 1
	}
}

func fibonacci(n int) (result int) {
	if n < 2 {
		return n
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
		return result
	}
}
