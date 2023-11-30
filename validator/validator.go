package validator

import (
	"fmt"
	"regexp"
)

const (
	isStatusRegexSpaceString = "[0-1]"
)

var (
	isStatusSpaceRegex = regexp.MustCompile(isStatusRegexSpaceString)
)

var Validate *validator.Validate

func Init() {
	Validate = validator.New()
	Validate.RegisterValidation("is_status", IsStatusSpace)
	Validate.RegisterValidation("option", OptionValidator)
}

func IsStatusSpace(field validator.FieldLevel) bool {
	number := field.Field().Int()

	return isStatusSpaceRegex.MatchString(fmt.Sprintf("%d", number))
}
