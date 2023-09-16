package main

import "fmt"

func main() {
	fmt.Println(fact(6))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(12))
}

func fact(x int) int {
	if x <= 1 {
		return 1
	}
	return x * fact(x-1)
}
