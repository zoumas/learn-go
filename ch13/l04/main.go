package main

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, e := range emails {
		ch <- e
	}
	return ch
}
