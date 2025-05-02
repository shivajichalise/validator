package rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/shivajichalise/validator"
)

// LtRule validates that a numeric field is less than a specified threshold.
// It supports int, float32, and float64 types.
type LtRule struct{}

func init() {
	validator.RegisterRule(LtRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "lt").
func (r LtRule) Name() string {
	return "lt"
}

// Validate checks whether the given numeric value is strictly less than the provided threshold.
// The threshold must be passed as a parameter (e.g., "lt:100").
// Returns an error if the value is not numeric, the parameter is missing,
// or if the comparison fails.
// If the value is an integer, the threshold must be a whole number.
func (r LtRule) Validate(field string, value any, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("%s: lt rule requires a comparison value", field)
	}

	ltValue, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return fmt.Errorf("%s: lt parameter must be a valid number", field)
	}

	num, err := validator.ToFloat64(value)
	if err != nil {
		return fmt.Errorf("%s must be numeric to use lt (apply 'numeric', 'int', or 'float64' rule first)", field)
	}

	if reflect.TypeOf(value).Kind() == reflect.Int && !validator.IsWholeNumber(ltValue) {
		return fmt.Errorf("lt value %.2f must be a whole number when %s is an integer", ltValue, field)
	}

	if num >= ltValue {
		return fmt.Errorf("%s must be less than %v", field, ltValue)
	}

	return nil
}
