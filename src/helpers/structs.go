package helpers

import (
	"reflect"
	"strings"
)

// ToMap converts a struct (including embedded structs) into a map[string]any.
// It uses reflection to handle any struct type.
func StructToMap(s interface{}) map[string]any {
	result := make(map[string]any)

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Dereference the pointer to get the underlying value
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		if fieldValue.CanInterface() {
			// If the field is an embedded struct, we recursively convert it to a map
			// but since we want all fields in the top-level map, we add each sub-field directly.
			if field.Anonymous && fieldValue.Kind() == reflect.Struct {
				embeddedMap := StructToMap(fieldValue.Addr().Interface())
				for k, v := range embeddedMap {
					result[strings.ToLower(k)] = v
				}
			} else {
				result[strings.ToLower(field.Name)] = fieldValue.Interface()
			}
		}
	}

	return result
}
