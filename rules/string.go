package rules

import (
	"fmt"
	"strings"

	"github.com/shivajichalise/validator"
)

// StringRule validates that a value is a non-empty string.
// It trims whitespace and considers an empty or whitespace-only string as invalid.
type StringRule struct{}

func init() {
	validator.RegisterRule(StringRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "string").
func (r StringRule) Name() string {
	return "string"
}

// Validate checks whether the given value is a string and not empty after trimming whitespace.
// Returns an error if the value is not a string or is empty.
func (r StringRule) Validate(field string, value any, _ ...string) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s field must be a valid string", field)
	}

	if strings.TrimSpace(str) == "" {
		return fmt.Errorf("%s must not be empty", field)
	}

	return nil
}
