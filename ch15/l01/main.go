package main

func getLast[T any](s []T) T {
	l := len(s)
	if l == 0 {
		var zero T
		return zero
	}
	return s[l-1]
}
