package rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/shivajichalise/validator"
)

// GtRule validates that a numeric field is greater than a specified threshold.
type GtRule struct{}

func init() {
	validator.RegisterRule(GtRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "gt").
func (r GtRule) Name() string {
	return "gt"
}

// Validate checks whether the given numeric value is strictly greater than the specified threshold.
// The threshold must be passed as a parameter (e.g., "gt:10").
// Returns an error if the value is not numeric, if the parameter is missing,
// or if the value is not greater than the threshold.
// If an integer is being compared, the threshold must be a whole number.
func (r GtRule) Validate(field string, value any, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("%s: gt rule requires a comparison value", field)
	}

	gtValue, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return fmt.Errorf("%s: gt parameter must be a valid number", field)
	}

	num, err := validator.ToFloat64(value)
	if err != nil {
		return fmt.Errorf("%s must be numeric to use gt (apply 'numeric', 'int', or 'float64' rule first)", field)
	}

	if reflect.TypeOf(value).Kind() == reflect.Int && !validator.IsWholeNumber(gtValue) {
		return fmt.Errorf("gt value %.2f must be a whole number when %s is an integer", gtValue, field)
	}

	if num <= gtValue {
		return fmt.Errorf("%s must be greater than %v", field, gtValue)
	}

	return nil
}
