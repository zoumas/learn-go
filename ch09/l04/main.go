package main

func getMessageCosts(messages []string) []float64 {
	l := len(messages)
	costs := make([]float64, l)
	for i := 0; i < l; i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}
