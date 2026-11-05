package main

func isValidPassword(password string) bool {
	l := len(password)
	if l < 5 || l > 12 {
		return false
	}

	var hasUpper, hasDigit bool
	for _, c := range password {
		if c >= '0' && c <= '9' {
			hasDigit = true
		} else if c >= 'A' && c <= 'Z' {
			hasUpper = true
		}
	}

	return hasDigit && hasUpper
}
