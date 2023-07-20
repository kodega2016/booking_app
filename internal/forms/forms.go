package forms

import (
	"net/http"
	"net/url"
)

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

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

// Has Required checks if a required field is in the form data and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}
