package main

import (
	"fmt"

	"github.com/shivajichalise/validator"
	_ "github.com/shivajichalise/validator/rules"
)

func main() {
	data := map[string]any{
		"name":               1,
		"username":           "",
		"email":              "rick@astley.com",
		"email_second":       "rick@test.com",
		"description":        "nevergonnagiveyouup",
		"description_second": "nevergonnaletyoudown",
	}

	rules := map[string][]string{
		"name":               {"string"},
		"username":           {"string"},
		"email":              {"email:rfc,dns"},
		"email_second":       {"email:rfc,dns"},
		"description":        {"string", "min:19"},
		"description_second": {"string", "min:5", "max:19"},
	}

	v := validator.Make(data, rules)

	if v.Validate() {
		fmt.Println("validation passed")
	} else {
		fmt.Println("validation failed")

		for field, messages := range v.Errors() {
			for _, message := range messages {
				fmt.Printf("'%s': %s\n", field, message)
			}
		}
	}
}
