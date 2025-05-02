package validator

import (
	"fmt"
	"reflect"
)

// IsWholeNumber returns true if the given float64 represents a whole number.
// For example, 5.0 returns true, while 5.3 returns false.
func IsWholeNumber(f float64) bool {
	return f == float64(int64(f))
}

// ToFloat64 attempts to convert supported numeric types to float64.
// Accepts int, int8, int16, int32, int64, float32, and float64.
// Returns an error if the value is not a supported numeric type.
func ToFloat64(value any) (float64, error) {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	default:
		return 0, fmt.Errorf("value of type %T is not a supported numeric type", value)
	}
}
