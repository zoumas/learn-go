package main

import "strings"

func countDistinctWords(messages []string) int {
	m := make(map[string]struct{})
	for _, msg := range messages {
		for s := range strings.FieldsSeq(strings.ToLower(msg)) {
			m[s] = struct{}{}
		}
	}
	return len(m)
}
