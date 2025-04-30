package rules

import (
	"fmt"
	"net"
	"net/mail"
	"regexp"
	"strings"

	"github.com/shivajichalise/validator"
)

type EmailRule struct{}

type emailValidationMode struct {
	basicOnly bool
	checkRFC  bool
	checkDNS  bool
}

var basicEmailRegex *regexp.Regexp

func init() {
	validator.RegisterRule(EmailRule{})

	basicEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
}

func (r EmailRule) Name() string {
	return "email"
}

func (r EmailRule) Validate(field string, value any, params ...string) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s field must be a valid string", field)
	}

	if strings.TrimSpace(str) == "" {
		return fmt.Errorf("%s must not be empty", field)
	}

	mode := parseEmailMode(params)

	// 1. just check the syntax for email
	if mode.basicOnly {
		if !basicEmailRegex.MatchString(str) {
			return fmt.Errorf("%s must be a valid email format (missing '@' or domain)", field)
		}
		return nil
	}

	addr, err := mail.ParseAddress(str)

	// 2. RFC compliant check
	if mode.checkRFC {
		if err != nil {
			return fmt.Errorf("%s must be a valid RFC-compliant email address", field)
		}
	}

	// 3. DNS MX check
	if mode.checkDNS {
		domain := strings.ToLower(strings.SplitN(addr.Address, "@", 2)[1])

		mxRecords, err := net.LookupMX(domain)
		if err != nil || len(mxRecords) == 0 {
			return fmt.Errorf("%s domain '%s' does not have valid MX records", field, domain)
		}
	}

	return nil
}

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
