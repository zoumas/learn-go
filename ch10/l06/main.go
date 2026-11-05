package main

func getNameCounts(names []string) map[rune]map[string]int {
	m := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		if len(runes) == 0 {
			continue
		}
		first := runes[0]
		if _, ok := m[first]; !ok {
			m[first] = make(map[string]int)
		}
		m[first][name]++
	}
	return m
}
