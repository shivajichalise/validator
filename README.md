# Validator

![Validator Banner](https://shivajichalise/validator/banner.png)

A lightweight, expressive, and extensible validation library for Go — inspired by Laravel’s Validator.

---

## Features

- Rule-based validation like `string`, `numeric`, `email`, `min`, `max`, `gt`, `lt`, etc.
- Chainable and expressive rule definitions
- Easy to extend with custom rules

---

## Installation

```bash
go get github.com/shivajichalise/validator
```

---

## Basic Usage

```go
import (
    "github.com/shivajichalise/validator"
    _ "github.com/shivajichalise/validator/rules"
)

data := map[string]any{
    "username": "rickastley",
    "email":    "rick@astley.com",
    "age":      21,
}

rules := map[string][]string{
    "username": {"string", "min:5", "max:20"},
    "email":    {"email:rfc,dns"},
    "age":      {"numeric", "gt:18"},
}

v := validator.Make(data, rules)

if !v.Validate() {
    for field, errs := range v.Errors() {
        for _, msg := range errs {
            fmt.Printf("%s: %s\n", field, msg)
        }
    }
}
```

---

## Supported Rules

| Rule      | Description                                   |
| --------- | --------------------------------------------- |
| `string`  | Value must be a non-empty string              |
| `min:n`   | String length or numeric value must be ≥ n    |
| `max:n`   | String length or numeric value must be ≤ n    |
| `email`   | Validates email with basic, RFC, or DNS check |
| `numeric` | Accepts int and float                         |
| `int`     | Value must be an integer                      |
| `float64` | Value must be a float64                       |
| `gt:n`    | Value must be greater than n                  |
| `lt:n`    | Value must be less than n                     |

---

## Rule Examples

```go
"username": {"string", "min:5", "max:20"}
"email": {"email:rfc,dns"}
"age": {"numeric", "gt:18"}
"score": {"int", "lt:100"}
```

---

## Custom Rules

Register a custom rule using:

```go
validator.RegisterRule(MyCustomRule{})
```

Your rule should implement:

```go
type Rule interface {
    Name() string
    Validate(field string, value any, params ...string) error
}
```

---

## Tests

This package is well-tested and includes:

- `string`, `min`, `max`
- `email` with basic, RFC, DNS
- `numeric`, `int`, `float64`
- `gt`, `lt` including type strictness and validation chaining
- Missing and invalid parameter handling

Run all tests:

```bash
go test ./...
```

---

## Contributing

We welcome contributions from the community! If you have an idea for a new feature or improvement, please submit a pull request. We also appreciate bug reports and other feedback.

To get started with contributing, simply fork this repository, make your changes, and submit a pull request.

## License

This project is licensed under [MIT](https://github.com/shivajichalise/validator/blob/main/LICENSE)

## Self-Promotion

Star the repository on [Github](https://github.com/shivajichalise/validator)

Follow on [Github](https://github.com/shivajichalise)
