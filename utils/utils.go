package utils

import (
	"regexp"
)

/*
* ValidateEmail checks if the email is valid
 */
func ValidateEmail(email string) bool {
	// Simple email validation
	if email == "" {
		return false
	}

	// Validate if is a real email format
	// This is a very basic regex for demonstration purposes
	if !regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`).MatchString(email) {
		return false
	}

	return true
}
