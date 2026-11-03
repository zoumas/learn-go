package main

import (
	"errors"
	"fmt"
)

func getLogger(label string) func(string) {
	return func(message string) {
		fmt.Println(formatLog(label, message))
	}
}

// don't touch below this line

func test(label string, errors []error) {
	defer fmt.Println("====================================")
	logger := getLogger(label)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(err.Error())
	}
}

func formatLog(label, message string) string {
	return label + ": " + message
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test("Error on mail server", mailErrors)
}
