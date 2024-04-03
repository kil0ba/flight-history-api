package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func Username(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	usernameRegex := `^[a-zA-Z][a-zA-Z0-9_-]*$`
	return matchRegex(usernameRegex, username)
}

// matchRegex checks if the given string matches the specified regular expression
func matchRegex(regex, str string) bool {
	match, _ := regexp.MatchString(regex, str)
	return match
}
