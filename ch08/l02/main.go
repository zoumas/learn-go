package main

func maxMessages(thresh int) int {
	var acc int
	for i := 0; ; i++ {
		acc += 100 + i
		if acc > thresh {
			return i
		}
	}
}
