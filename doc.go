// Package validator provides a lightweight, expressive, and extensible validation library for Go,
// inspired by Laravel's Validator. It allows developers to define validation rules using
// declarative, chainable syntax on dynamic input data.
//
// This package supports a wide range of rules such as:
//
//   - string, min, max
//   - email (basic, rfc, dns)
//   - numeric, int, float64
//   - gt, lt (greater/less than)
//   - boolean
//
// Example usage:
//
//	data := map[string]any{
//	    "email": "rick@astley.com",
//	    "age":   21,
//	}
//
//	rules := map[string][]string{
//	    "email": {"email:rfc,dns"},
//	    "age":   {"numeric", "gt:18"},
//	}
//
//	v := validator.Make(data, rules)
//	if !v.Validate() {
//	    for field, errs := range v.Errors() {
//	        for _, msg := range errs {
//	            fmt.Println(field + ": " + msg)
//	        }
//	    }
//	}
//
// See README for full examples, available rules, and custom rule extension.
package validator
