package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, m := range msg {
		for _, bad := range badWords {
			if m == bad {
				return i
			}
		}
	}
	return -1
}
