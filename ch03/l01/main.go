package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen > maxMessageLen {
		fmt.Println("Message not sent")
	} else {
		fmt.Println("Message sent")
	}
}
