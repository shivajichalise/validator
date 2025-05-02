package rules

import (
	"fmt"
	"strconv"

	"github.com/shivajichalise/validator"
)

// MaxRule validates that the length of a string does not exceed a specified maximum.
// It only applies to string values.
// Use "max:n" in your rule expression, where n is the maximum allowed length.
type MaxRule struct{}

func init() {
	validator.RegisterRule(MaxRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "max").
func (r MaxRule) Name() string {
	return "max"
}

// Validate checks whether the length of a string value is less than or equal to the given maximum.
// The maximum length must be provided as a parameter (e.g., "max:10").
// Returns an error if the value is not a string, the parameter is missing, or the string exceeds the maximum length.
func (r MaxRule) Validate(field string, value any, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("%s: max rule requires a length parameter", field)
	}

	maxLen, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("%s: max value must be a valid number", field)
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string to use max", field)
	}

	if len(str) > maxLen {
		return fmt.Errorf("%s must be at most %d characters", field, maxLen)
	}

	return nil
}
