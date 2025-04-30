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
