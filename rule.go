package validator

type Rule interface {
	Name() string
	Validate(field string, value any) error
}
