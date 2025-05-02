package rules

import (
	"fmt"
	"strconv"

	"github.com/shivajichalise/validator"
)

// MinRule validates that the length of a string is at least a specified minimum.
// This rule applies only to string values.
// Use "min:n" in rule expressions, where n is the required minimum number of characters.
type MinRule struct{}

func init() {
	validator.RegisterRule(MinRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "min").
func (r MinRule) Name() string {
	return "min"
}

// Validate checks whether the length of a string value is greater than or equal to the specified minimum.
// The minimum length must be provided as a parameter (e.g., "min:5").
// Returns an error if the value is not a string, the parameter is missing, or the string is too short.
func (r MinRule) Validate(field string, value any, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("%s: min rule requires a length parameter", field)
	}

	minLen, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("%s: min value must be a valid number", field)
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string to use min", field)
	}

	if len(str) < minLen {
		return fmt.Errorf("%s must be at least %d characters", field, minLen)
	}

	return nil
}
