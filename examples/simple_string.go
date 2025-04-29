package main

import (
	"fmt"

	"github.com/shivajichalise/validator"
	_ "github.com/shivajichalise/validator/rules"
)

func main() {
	data := map[string]any{
		"name":        1,
		"username":    "",
		"description": "nevergonnagiveyouup",
	}

	rules := map[string][]string{
		"name":        {"string"},
		"username":    {"string"},
		"description": {"string"},
	}

	v := validator.Make(data, rules)

	if v.Validate() {
		fmt.Println("valid username")
	} else {
		fmt.Println("invalid username")

		for field, messages := range v.Errors() {
			for _, message := range messages {
				fmt.Printf("'%s': %s\n", field, message)
			}
		}
	}
}
