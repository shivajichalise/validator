package rules

import (
	"fmt"

	"github.com/shivajichalise/validator"
)

type BooleanRule struct{}

func init() {
	validator.RegisterRule(BooleanRule{})
}

func (r BooleanRule) Name() string {
	return "boolean"
}

func (r BooleanRule) Validate(field string, value any, _ ...string) error {
	switch v := value.(type) {
	case bool:
		return nil
	case string:
		if v == "true" || v == "false" || v == "1" || v == "0" {
			return nil
		}
	case int:
		if v == 0 || v == 1 {
			return nil
		}
	}

	return fmt.Errorf("%s must be a boolean value (true, false, 1, 0)", field)
}
