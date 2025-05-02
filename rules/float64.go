package rules

import (
	"fmt"
	"reflect"

	"github.com/shivajichalise/validator"
)

// Float64Rule ensures that a field's value is strictly of type float64.
// This rule enforces type safety in validation chains that depend on numeric precision.
type Float64Rule struct{}

func init() {
	validator.RegisterRule(Float64Rule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "float64").
func (r Float64Rule) Name() string {
	return "float64"
}

// Validate checks whether the given value is of type float64.
// Returns an error if the value is not exactly a float64.
func (r Float64Rule) Validate(field string, value any, _ ...string) error {
	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		return nil
	}
	return fmt.Errorf("%s must be a float64 value", field)
}
