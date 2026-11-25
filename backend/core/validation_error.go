package core

var _ error = ValidationError("")

type ValidationError string

func (s ValidationError) Error() string {
	return string(s)
}

func IsValidationError(err any) bool {
	if _, ok := err.(ValidationError); ok {
		return true
	}
	return false
}
