package main

func bulkSend(numMessages int) float64 {
	var totalCost float64

	for i := range numMessages {
		totalCost += 1.0 + float64(i)*0.01
	}

	return totalCost
}
