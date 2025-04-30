package rules

import (
	"fmt"
	"strconv"

	"github.com/shivajichalise/validator"
)

type MaxRule struct{}

func init() {
	validator.RegisterRule(MaxRule{})
}

func (r MaxRule) Name() string {
	return "max"
}

func (r MaxRule) Validate(field string, value any, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("%s: max rule requires a length parameter", field)
	}

	maxLen, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("%s: max value must be a valid number", field)
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string to use max", field)
	}

	if len(str) > maxLen {
		return fmt.Errorf("%s must be at most %d characters", field, maxLen)
	}

	return nil
}
