package rules

import (
	"fmt"

	"github.com/shivajichalise/validator"
)

// NumericRule validates that a value is numeric.
// It accepts all standard Go numeric types: int, int8, int16, int32, int64, float32, and float64.
type NumericRule struct{}

func init() {
	validator.RegisterRule(NumericRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "numeric").
func (r NumericRule) Name() string {
	return "numeric"
}

// Validate checks whether the given value is a supported numeric type.
// Returns an error if the value is not numeric.
// Internally uses validator.ToFloat64 to normalize numeric types.
func (r NumericRule) Validate(field string, value any, _ ...string) error {
	_, err := validator.ToFloat64(value)
	if err != nil {
		return fmt.Errorf("%s must be a numeric value", field)
	}
	return nil
}
