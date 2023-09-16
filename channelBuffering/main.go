package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "bufferred"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
