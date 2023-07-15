package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom Form struct, embeds a url.Values object, and adds an errors field
type Form struct {
	url.Values
	Errors errors
}

// New initializes a custom Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if a required field is in the form data and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
