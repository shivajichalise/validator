package rules

import (
	"fmt"
	"strconv"

	"github.com/shivajichalise/validator"
)

type MinRule struct{}

func init() {
	validator.RegisterRule(MinRule{})
}

func (r MinRule) Name() string {
	return "min"
}

func (r MinRule) Validate(field string, value any, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("%s: min rule requires a length parameter", field)
	}

	minLen, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("%s: min value must be a valid number", field)
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string to use min", field)
	}

	if len(str) < minLen {
		return fmt.Errorf("%s must be at least %d characters", field, minLen)
	}

	return nil
}
