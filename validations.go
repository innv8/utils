package utils


import (
	"regexp"
)

// ValidateUserName
// validates if a username is of correct format
// small letters and/numbers numbers
// must be 5 characters or more
func ValidateUserName(userName string) bool {
	match, _ := regexp.MatchString("[a-z0-9]{5,}", userName)
	return match
}

// ValidateEmail
// validates if an email is valid
func ValidateEmail(email string) bool {
	match, _ := regexp.MatchString("[a-zA-Z0-9._]{2,}@[a-z]{1,}[.a-z]{1,}", email)
	return match
}

// ValidatePassword
// validates a password, can have letters and numbers
// must be 6 characters or more
func ValidatePassword(password string) bool {
	match, _ := regexp.MatchString("[a-zA-Z0-9]{6,}", password)
	return match
}