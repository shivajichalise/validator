package rules

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/shivajichalise/validator"
)

// BetweenRule validates that a numeric field lies strictly between two specified values.
type BetweenRule struct{}

func init() {
	validator.RegisterRule(BetweenRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "between").
func (r BetweenRule) Name() string {
	return "between"
}

// Validate checks whether the given numeric value lies strictly between two thresholds.
// The thresholds must be passed as a single comma-separated parameter (e.g., "between:1,10").
// Returns an error if:
// - the parameter is missing or incorrectly formatted
// - the value is not numeric
// - the value is not strictly between the provided thresholds
// - the field is an integer and thresholds are not whole numbers
func (r BetweenRule) Validate(field string, value any, params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("%s: between rule requires a single parameter in the format 'min,max'", field)
	}

	parts := strings.Split(params[0], ",")
	if len(parts) != 2 {
		return fmt.Errorf("%s: between rule requires two comma-separated values", field)
	}

	min, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return fmt.Errorf("%s: lower cap must be a valid number", field)
	}

	max, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return fmt.Errorf("%s: upper cap must be a valid number", field)
	}

	val, err := validator.ToFloat64(value)
	if err != nil {
		return fmt.Errorf("%s must be numeric to use between (apply 'numeric', 'int', or 'float64' rule first)", field)
	}

	kind := reflect.TypeOf(value).Kind()
	if kind == reflect.Int {
		if !validator.IsWholeNumber(min) {
			return fmt.Errorf("between value %.2f must be a whole number when %s is an integer", min, field)
		}
		if !validator.IsWholeNumber(max) {
			return fmt.Errorf("between value %.2f must be a whole number when %s is an integer", max, field)
		}
	}

	if !(min < val && val < max) {
		return fmt.Errorf("%s must be between %v and %v", field, min, max)
	}

	return nil
}
