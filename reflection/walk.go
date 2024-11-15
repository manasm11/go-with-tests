package walk

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

  if val.Kind() == reflect.Map {
    for _, key := range val.MapKeys() {
      walk(val.MapIndex(key).Interface(), fn)
    }
  }

	var numberOfValues int
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		getField = val.Field
		numberOfValues = val.NumField()
	case reflect.Slice, reflect.Array:
		getField = val.Index
		numberOfValues = val.Len()
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
