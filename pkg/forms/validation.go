package forms

import "strings"

func IsEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}
