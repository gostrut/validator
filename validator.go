package validator

import "reflect"
import "github.com/gostrut/invalid"

// Func is the signature for a Validator function
type Func func(string, string, *reflect.Value) (invalid.Field, error)

// Validate calls v(field, tagStr, vo)
func (v Func) Validate(field, tagStr string, vo *reflect.Value) (invalid.Field, error) {
	return v(field, tagStr, vo)
}
