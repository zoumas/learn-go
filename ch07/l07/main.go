package main

import (
	"errors"
)

func validateStatus(status string) error {
	l := len(status)
	if l == 0 {
		return errors.New("status cannot be empty")
	}
	if l > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}
