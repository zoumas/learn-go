package main

func sum(nums ...int) int {
	var acc int
	for i := 0; i < len(nums); i++ {
		acc += nums[i]
	}
	return acc
}
