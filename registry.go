package validator

import (
	"fmt"
)

// ruleRegistry holds all registered validation rules by their name.
var ruleRegistry = make(map[string]Rule)

// RegisterRule adds a new rule implementation to the global registry.
// It panics if a rule with the same name has already been registered.
// Typically called in init() functions inside rule packages.
func RegisterRule(rule Rule) {
	name := rule.Name()

	_, exists := ruleRegistry[name]
	if exists {
		panic(fmt.Sprintf("rule '%s' is already registered", name))
	}

	ruleRegistry[rule.Name()] = rule
}

// GetRule retrieves a rule implementation by its name.
// Returns the rule and true if found, otherwise returns false.
func GetRule(name string) (Rule, bool) {
	rule, ok := ruleRegistry[name]

	return rule, ok
}
