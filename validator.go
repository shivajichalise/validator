package validator

import "fmt"

type (
	data   map[string]any
	rules  map[string][]string
	Errors map[string][]string
)

type Validator struct {
	data   data
	rules  rules
	errors Errors
}

func Make(data data, rules rules) *Validator {
	return &Validator{
		data:   data,
		rules:  rules,
		errors: make(Errors),
	}
}

func (v *Validator) Errors() Errors {
	return v.errors
}

func (v *Validator) addError(field, message string) {
	v.errors[field] = append(v.errors[field], message)
}

func parseRule(expr string) (string, []string) {
	for i, c := range expr {
		if c == ':' {
			return expr[:i], []string{expr[i+1:]}
		}
	}
	return expr, nil
}

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
