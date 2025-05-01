package rules

import (
	"fmt"
	"reflect"

	"github.com/shivajichalise/validator"
)

type Float64Rule struct{}

func init() {
	validator.RegisterRule(Float64Rule{})
}

func (r Float64Rule) Name() string {
	return "float64"
}

func (r Float64Rule) Validate(field string, value any, _ ...string) error {
	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		return nil
	}
	return fmt.Errorf("%s must be a float64 value", field)
}
