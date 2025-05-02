package rules

import (
	"fmt"
	"reflect"

	"github.com/shivajichalise/validator"
)

// IntRule ensures that a field's value is an integer type.
// Supported types include int, int8, int16, int32, and int64.
type IntRule struct{}

func init() {
	validator.RegisterRule(IntRule{})
}

// Name returns the name of the rule used in rule expressions (e.g., "int").
func (r IntRule) Name() string {
	return "int"
}

// Validate checks whether the value is of an integer type.
// Returns an error if the value is not a supported integer kind.
func (r IntRule) Validate(field string, value any, _ ...string) error {
	kind := reflect.TypeOf(value).Kind()

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil
	default:
		return fmt.Errorf("%s must be an integer", field)
	}
}
