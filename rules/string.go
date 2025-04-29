package rules

import (
	"fmt"
	"strings"

	"github.com/shivajichalise/validator"
)

type StringRule struct{}

func init() {
	validator.RegisterRule(StringRule{})
}

func (r StringRule) Name() string {
	return "string"
}

func (r StringRule) Validate(field string, value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s field must be a valid string", field)
	}

	if strings.TrimSpace(str) == "" {
		return fmt.Errorf("%s must not be empty", field)
	}

	return nil
}
