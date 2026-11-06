package main

import (
	"strings"
)

func removeProfanity(message *string) {
	s := *message

	s = strings.ReplaceAll(s, "fubb", "****")
	s = strings.ReplaceAll(s, "shiz", "****")
	s = strings.ReplaceAll(s, "witch", "*****")

	*message = s
}
