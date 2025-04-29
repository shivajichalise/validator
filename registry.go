package validator

import (
	"fmt"
)

var ruleRegistry = make(map[string]Rule)

func RegisterRule(rule Rule) {
	name := rule.Name()

	_, exists := ruleRegistry[name]
	if exists {
		panic(fmt.Sprintf("rule '%s' is already registered", name))
	}

	ruleRegistry[rule.Name()] = rule
}

func GetRule(name string) (Rule, bool) {
	rule, ok := ruleRegistry[name]

	return rule, ok
}
