package validator

// Rule defines the interface for all validation rules.
// Each rule must have a unique name and implement custom validation logic.
type Rule interface {
	// Name returns the identifier for the rule (e.g., "string", "min", "email").
	// This name is used when mapping rule strings to their implementations.
	Name() string

	// Validate runs the validation logic for the given field and value.
	// The params slice contains any optional arguments passed with the rule
	// (e.g., "min:5" would pass "5" as params[0]).
	// It returns an error if the validation fails.
	Validate(field string, value any, params ...string) error
}
