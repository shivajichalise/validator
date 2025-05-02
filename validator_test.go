// Package validator_test provides table-driven tests for all built-in validation rules
// in the validator package.
//
// Each test case defines input data, associated rules, and the expected outcome.
// These tests verify rule correctness, edge-case handling, and type enforcement
// across string, numeric, email, boolean, and comparison validations.
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
		{
			name: "int value with int rule",
			data: map[string]any{"age": 42},
			rules: map[string][]string{
				"age": {"int"},
			},
			wantErr: false,
		},
		{
			name: "float64 value with int rule",
			data: map[string]any{"age": 42.0},
			rules: map[string][]string{
				"age": {"int"},
			},
			wantErr: true,
		},
		{
			name: "float64 value with float64 rule",
			data: map[string]any{"rating": 4.2},
			rules: map[string][]string{
				"rating": {"float64"},
			},
			wantErr: false,
		},
		{
			name: "int value with float64 rule",
			data: map[string]any{"rating": 4},
			rules: map[string][]string{
				"rating": {"float64"},
			},
			wantErr: true,
		},
		{
			name: "int value with numeric rule",
			data: map[string]any{"score": 100},
			rules: map[string][]string{
				"score": {"numeric"},
			},
			wantErr: false,
		},
		{
			name: "float value with numeric rule",
			data: map[string]any{"score": 99.9},
			rules: map[string][]string{
				"score": {"numeric"},
			},
			wantErr: false,
		},
		{
			name: "string value with numeric rule",
			data: map[string]any{"score": "high"},
			rules: map[string][]string{
				"score": {"numeric"},
			},
			wantErr: true,
		},
		{
			name: "gt with numeric int value passes",
			data: map[string]any{"amount": 10},
			rules: map[string][]string{
				"amount": {"numeric", "gt:5"},
			},
			wantErr: false,
		},
		{
			name: "gt with numeric float value passes",
			data: map[string]any{"amount": 10.5},
			rules: map[string][]string{
				"amount": {"numeric", "gt:10.1"},
			},
			wantErr: false,
		},
		{
			name: "gt with value equal to threshold fails",
			data: map[string]any{"amount": 5},
			rules: map[string][]string{
				"amount": {"numeric", "gt:5"},
			},
			wantErr: true,
		},
		{
			name: "gt with float threshold and int type rule fails",
			data: map[string]any{"amount": 16},
			rules: map[string][]string{
				"amount": {"int", "gt:16.1"},
			},
			wantErr: true,
		},
		{
			name: "gt with missing parameter",
			data: map[string]any{"amount": 20},
			rules: map[string][]string{
				"amount": {"numeric", "gt"},
			},
			wantErr: true,
		},
		{
			name: "lt with int less than threshold",
			data: map[string]any{"score": 8},
			rules: map[string][]string{
				"score": {"numeric", "lt:10"},
			},
			wantErr: false,
		},
		{
			name: "lt with int equal to threshold",
			data: map[string]any{"score": 10},
			rules: map[string][]string{
				"score": {"numeric", "lt:10"},
			},
			wantErr: true,
		},
		{
			name: "lt with int greater than threshold",
			data: map[string]any{"score": 12},
			rules: map[string][]string{
				"score": {"numeric", "lt:10"},
			},
			wantErr: true,
		},
		{
			name: "lt with float less than threshold",
			data: map[string]any{"score": 9.99},
			rules: map[string][]string{
				"score": {"numeric", "lt:10.0"},
			},
			wantErr: false,
		},
		{
			name: "lt with float equal to threshold",
			data: map[string]any{"score": 10.0},
			rules: map[string][]string{
				"score": {"numeric", "lt:10.0"},
			},
			wantErr: true,
		},
		{
			name: "lt with int and float threshold mismatch",
			data: map[string]any{"score": 5},
			rules: map[string][]string{
				"score": {"int", "lt:5.5"},
			},
			wantErr: true,
		},
		{
			name: "lt with missing param",
			data: map[string]any{"score": 3},
			rules: map[string][]string{
				"score": {"numeric", "lt"},
			},
			wantErr: true,
		},
		{
			name: "valid boolean true",
			data: map[string]any{"is_active": true},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "valid boolean false",
			data: map[string]any{"is_active": false},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "valid string 'true'",
			data: map[string]any{"is_active": "true"},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "valid string 'false'",
			data: map[string]any{"is_active": "false"},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "valid string '1'",
			data: map[string]any{"is_active": "1"},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "valid string '0'",
			data: map[string]any{"is_active": "0"},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "valid int 1",
			data: map[string]any{"is_active": 1},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "valid int 0",
			data: map[string]any{"is_active": 0},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: false,
		},
		{
			name: "invalid string value",
			data: map[string]any{"is_active": "yes"},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: true,
		},
		{
			name: "invalid int value",
			data: map[string]any{"is_active": 2},
			rules: map[string][]string{
				"is_active": {"boolean"},
			},
			wantErr: true,
		},
		{
			name: "nil value fails",
			data: map[string]any{"is_active": nil},
			rules: map[string][]string{
				"is_active": {"boolean"},
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
