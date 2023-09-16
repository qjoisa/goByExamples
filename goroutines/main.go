package main

import (
	"fmt"
	"time"
)

func main() {
	go f("direct")

	go f("goroutine")
	time.Sleep(time.Millisecond)
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
func f(from string) {
	for i := 0; i < 1000000; i++ {
		fmt.Println(from, ":", i)
	}
}
