package rules

import (
	"fmt"
	"net"
	"net/mail"
	"regexp"
	"strings"

	"github.com/shivajichalise/validator"
)

// EmailRule validates whether a value is a properly formatted email address.
// It supports multiple modes: basic format check, RFC-compliant syntax, and DNS MX lookup.
type EmailRule struct{}

// emailValidationMode controls which levels of email validation are enabled.
type emailValidationMode struct {
	basicOnly bool // Perform only a simple format check
	checkRFC  bool // Enable RFC-compliant syntax validation
	checkDNS  bool // Perform MX record lookup for domain
}

// basicEmailRegex is used for simple format validation when no advanced checks are enabled.
var basicEmailRegex *regexp.Regexp

func init() {
	validator.RegisterRule(EmailRule{})

	basicEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
}

// Name returns the name of the rule used in rule expressions (e.g., "email").
func (r EmailRule) Name() string {
	return "email"
}

// Validate performs email validation on the given field value based on the selected mode.
// Supported modes (via params):
//   - "rfc": enables RFC-compliant syntax check
//   - "dns": enables MX record lookup on domain
//
// If no parameters are provided, only the basic format is validated.
func (r EmailRule) Validate(field string, value any, params ...string) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s field must be a valid string", field)
	}

	if strings.TrimSpace(str) == "" {
		return fmt.Errorf("%s must not be empty", field)
	}

	mode := parseEmailMode(params)

	// 1. Basic format validation using regex
	if mode.basicOnly {
		if !basicEmailRegex.MatchString(str) {
			return fmt.Errorf("%s must be a valid email format (missing '@' or domain)", field)
		}
		return nil
	}

	addr, err := mail.ParseAddress(str)

	// 2. RFC-compliant email validation
	if mode.checkRFC {
		if err != nil {
			return fmt.Errorf("%s must be a valid RFC-compliant email address", field)
		}
	}

	// 3. DNS MX record check on domain
	if mode.checkDNS {
		domain := strings.ToLower(strings.SplitN(addr.Address, "@", 2)[1])

		mxRecords, err := net.LookupMX(domain)
		if err != nil || len(mxRecords) == 0 {
			return fmt.Errorf("%s domain '%s' does not have valid MX records", field, domain)
		}
	}

	return nil
}

// parseEmailMode parses rule parameters and returns the enabled validation modes.
// Defaults to basic-only validation if no parameters are specified.
func parseEmailMode(params []string) emailValidationMode {
	if len(params) == 0 {
		return emailValidationMode{basicOnly: true}
	}

	mode := emailValidationMode{}
	flags := strings.Split(params[0], ",")

	for _, flag := range flags {
		switch strings.TrimSpace(strings.ToLower(flag)) {
		case "rfc":
			mode.checkRFC = true
		case "dns":
			mode.checkDNS = true
		}
	}

	return mode
}
