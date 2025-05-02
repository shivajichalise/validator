package rules

import (
	"fmt"

	"github.com/shivajichalise/validator"
)

// BooleanRule checks if a field is a boolean value.
// It accepts Go bool types, or common equivalents like "true", "false", "1", and "0" as strings or integers.
type BooleanRule struct{}

func init() {
	validator.RegisterRule(BooleanRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "boolean").
func (r BooleanRule) Name() string {
	return "boolean"
}

// Validate ensures the given value is a boolean or boolean-equivalent.
// Accepted values include:
//   - bool types: true, false
//   - strings: "true", "false", "1", "0"
//   - integers: 1, 0
//
// It returns an error if the value doesn't match any supported boolean form.
func (r BooleanRule) Validate(field string, value any, _ ...string) error {
	switch v := value.(type) {
	case bool:
		return nil
	case string:
		if v == "true" || v == "false" || v == "1" || v == "0" {
			return nil
		}
	case int:
		if v == 0 || v == 1 {
			return nil
		}
	}

	return fmt.Errorf("%s must be a boolean value (true, false, 1, 0)", field)
}
