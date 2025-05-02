package validator

import "fmt"

// data represents input values to be validated.
// Keys are field names, and values are the actual data.
type data map[string]any

// rules represents validation rules to be applied to each field.
// Keys are field names, and values are slices of rule expressions (e.g. "string", "min:5").
type rules map[string][]string

// Errors represents a map of validation errors.
// Keys are field names, and values are slices of error messages for that field.
type Errors map[string][]string

// Validator is the core struct that holds input data, validation rules, and error state.
type Validator struct {
	data   data
	rules  rules
	errors Errors
}

// Make creates a new Validator instance with the provided data and rules.
func Make(data data, rules rules) *Validator {
	return &Validator{
		data:   data,
		rules:  rules,
		errors: make(Errors),
	}
}

// Errors returns the collected validation errors after running Validate().
func (v *Validator) Errors() Errors {
	return v.errors
}

// addError appends an error message to the list of errors for a given field.
func (v *Validator) addError(field, message string) {
	v.errors[field] = append(v.errors[field], message)
}

// parseRule splits a rule expression into the rule name and its parameters.
// For example, "min:5" becomes ("min", ["5"]).
func parseRule(expr string) (string, []string) {
	for i, c := range expr {
		if c == ':' {
			return expr[:i], []string{expr[i+1:]}
		}
	}
	return expr, nil
}

// Validate runs all the rules against the data.
// It populates the internal errors map if any validations fail.
// Returns true if validation passes with no errors, false otherwise.
func (v *Validator) Validate() bool {
	for field, fieldRules := range v.rules {
		value := v.data[field]

		for _, ruleExpr := range fieldRules {
			ruleName, params := parseRule(ruleExpr)
			rule, exists := GetRule(ruleName)
			if !exists {
				v.addError(field, fmt.Sprintf("rule '%s' not found", ruleName))
				continue
			}

			err := rule.Validate(field, value, params...)
			if err != nil {
				v.addError(field, err.Error())
			}

		}
	}

	return len(v.errors) == 0
}
