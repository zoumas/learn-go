package main

func createMatrix(rows, cols int) [][]int {
	s := make([][]int, rows)
	for i := range rows {
		s[i] = make([]int, cols)
		for j := range cols {
			s[i][j] = i * j
		}
	}
	return s
}
