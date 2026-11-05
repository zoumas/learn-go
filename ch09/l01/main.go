package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	l1 := len(primary)
	l2 := len(secondary) + l1
	l3 := len(tertiary) + l2

	return [3]string{primary, secondary, tertiary}, [3]int{l1, l2, l3}
}
