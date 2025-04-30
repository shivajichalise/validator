package validator_test

import (
	"testing"

	"github.com/shivajichalise/validator"
	_ "github.com/shivajichalise/validator/rules"
)

func TestStringRule(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]any
		rules   map[string][]string
		wantErr bool
	}{
		{
			name: "valid string",
			data: map[string]any{"username": "rickastley"},
			rules: map[string][]string{
				"username": {"string"},
			},
			wantErr: false,
		},
		{
			name: "empty string",
			data: map[string]any{"username": ""},
			rules: map[string][]string{
				"username": {"string"},
			},
			wantErr: true,
		},
		{
			name: "non-string type",
			data: map[string]any{"username": 123},
			rules: map[string][]string{
				"username": {"string"},
			},
			wantErr: true,
		},
		{
			name: "string meets min",
			data: map[string]any{"username": "we're no strangers to love"},
			rules: map[string][]string{
				"username": {"string", "min:26"},
			},
			wantErr: false,
		},
		{
			name: "string below min",
			data: map[string]any{"username": "you know the rules and so do I"},
			rules: map[string][]string{
				"username": {"string", "min:31"},
			},
			wantErr: true,
		},
		{
			name: "min missing param",
			data: map[string]any{"username": "a full commitment's what I'm thinking of"},
			rules: map[string][]string{
				"username": {"string", "min"},
			},
			wantErr: true,
		},
		{
			name: "string meets max",
			data: map[string]any{"username": "you wouldn't get this from any other guy"},
			rules: map[string][]string{
				"username": {"string", "min:5", "max:40"},
			},
			wantErr: false,
		},
		{
			name: "string below min",
			data: map[string]any{"username": "I just wanna tell you how I'm feeling"},
			rules: map[string][]string{
				"username": {"string", "min:5", "max:10"},
			},
			wantErr: true,
		},
		{
			name: "max missing param",
			data: map[string]any{"username": "gotta make you understand"},
			rules: map[string][]string{
				"username": {"string", "max"},
			},
			wantErr: true,
		},
		{
			name: "basic email valid format",
			data: map[string]any{"email": "rick@astley.com"},
			rules: map[string][]string{
				"email": {"email"},
			},
			wantErr: false,
		},
		{
			name: "basic email missing domain",
			data: map[string]any{"email": "rick@astleycom"},
			rules: map[string][]string{
				"email": {"email"},
			},
			wantErr: true,
		},
		{
			name: "basic email missing @",
			data: map[string]any{"email": "rickastley.com"},
			rules: map[string][]string{
				"email": {"email"},
			},
			wantErr: true,
		},
		{
			name: "rfc email valid",
			data: map[string]any{"email": "rick@example.com"},
			rules: map[string][]string{
				"email": {"email:rfc"},
			},
			wantErr: false,
		},
		{
			name: "rfc email invalid format",
			data: map[string]any{"email": "rick@invalid..com"},
			rules: map[string][]string{
				"email": {"email:rfc"},
			},
			wantErr: true,
		},
		{
			name: "dns email valid MX",
			data: map[string]any{"email": "rick@google.com"},
			rules: map[string][]string{
				"email": {"email:dns"},
			},
			wantErr: false,
		},
		{
			name: "dns email invalid MX",
			data: map[string]any{"email": "rick@invalid.tld"},
			rules: map[string][]string{
				"email": {"email:dns"},
			},
			wantErr: true,
		},
		{
			name: "rfc and dns email valid",
			data: map[string]any{"email": "rick@gmail.com"},
			rules: map[string][]string{
				"email": {"email:rfc,dns"},
			},
			wantErr: false,
		},
		{
			name: "rfc and dns email invalid",
			data: map[string]any{"email": "rick@invalid.tld"},
			rules: map[string][]string{
				"email": {"email:rfc,dns"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.Make(tt.data, tt.rules)
			valid := v.Validate()

			if valid == tt.wantErr {
				t.Errorf(
					"\nTest: %s\nExpected valid: %v\nActual valid: %v\nValidation Errors:\n%v\n",
					tt.name,
					!tt.wantErr,
					valid,
					v.Errors(),
				)
			}
		})
	}
}
