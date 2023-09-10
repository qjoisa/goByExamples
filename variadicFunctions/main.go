package main

import "fmt"

func main() {
	sum(1, 3)
	sum(3, 5, 9)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
