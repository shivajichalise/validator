package rules

import (
	"fmt"
	"reflect"

	"github.com/shivajichalise/validator"
)

type IntRule struct{}

func init() {
	validator.RegisterRule(IntRule{})
}

func (r IntRule) Name() string {
	return "int"
}

func (r IntRule) Validate(field string, value any, _ ...string) error {
	kind := reflect.TypeOf(value).Kind()

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil
	default:
		return fmt.Errorf("%s must be an integer", field)
	}
}
