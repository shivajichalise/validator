package rules

import (
	"fmt"

	"github.com/shivajichalise/validator"
)

type NumericRule struct{}

func init() {
	validator.RegisterRule(NumericRule{})
}

func (r NumericRule) Name() string {
	return "numeric"
}

func (r NumericRule) Validate(field string, value any, _ ...string) error {
	_, err := validator.ToFloat64(value)
	if err != nil {
		return fmt.Errorf("%s must be a numeric value", field)
	}
	return nil
}
