package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, m := range messages {
		messages[i].tags = tagger(m)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	l := strings.ToLower(msg.content)
	if strings.Contains(l, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(l, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
