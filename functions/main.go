package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2", res)

	res = plusPlus(1, 3, 5)
	fmt.Println("1 + 2 + 3 =", res)
}
func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}
