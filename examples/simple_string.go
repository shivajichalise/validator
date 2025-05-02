package main

import (
	"fmt"

	"github.com/shivajichalise/validator"
	_ "github.com/shivajichalise/validator/rules"
)

func main() {
	data := map[string]any{
		"name":           123,
		"username":       "",
		"email":          "rick@astley.com",
		"backup_email":   "rick@test.com",
		"bio":            "nevergonnagiveyouup",
		"tagline":        "nevergonnaletyoudown",
		"age":            "17",
		"height":         17.5,
		"login_attempts": 17,
		"max_score":      80,
		"is_verified":    "1",
	}

	rules := map[string][]string{
		"name":           {"string"},                    // should fail: not a string
		"username":       {"string"},                    // should fail: empty
		"email":          {"email:rfc,dns"},             // should pass
		"backup_email":   {"email:rfc,dns"},             // should fail (invalid MX)
		"bio":            {"string", "min:19"},          // should pass
		"tagline":        {"string", "min:5", "max:19"}, // should pass
		"age":            {"numeric"},                   // should pass
		"height":         {"int"},                       // should fail: float given
		"login_attempts": {"float64"},                   // should fail: int given
		"max_score":      {"int", "lt:81"},              // should pass
		"is_verified":    {"boolean"},                   // should pass: string "1" allowed
	}

	v := validator.Make(data, rules)

	if v.Validate() {
		fmt.Println("All validations passed. Rick would be proud.")
	} else {
		fmt.Println("Validation failed. Someone let Rick down.")

		for field, messages := range v.Errors() {
			for _, message := range messages {
				fmt.Printf(" - %s: %s\n", field, message)
			}
		}
	}
}
