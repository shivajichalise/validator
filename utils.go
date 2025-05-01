package validator

import (
	"fmt"
	"reflect"
)

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
