package main

func countConnections(groupSize int) int {
	var c int
	for i := 1; i <= groupSize; i++ {
		c += groupSize - i
	}
	return c
}
