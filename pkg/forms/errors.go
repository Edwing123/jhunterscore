package forms

// This type will be used th hold form error messages.
type Errors map[string][]string

// Add an error message to the corresponding field.
func (e Errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

// Gets the last error of the slice, if any.
func (e Errors) Get(field string) string {
	fieldErrors := e[field]

	if len(fieldErrors) == 0 {
		return ""
	}

	return fieldErrors[0]
}
