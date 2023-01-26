package basic_benchmark

import "reflect"

func addInts(a, b int) int {
	return a + b
}
func addFloats(a, b float64) float64 {
	return a + b
}
func addStrings(a, b string) string {
	return a + b
}

func addWithTypeCastingUsingSwitch(a, b interface{}) (interface{}, bool) {
	switch a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a.(int) + b, true
		}
	case float64:
		if b, ok := b.(float64); ok {
			return a.(float64) + b, true
		}
	case string:
		if b, ok := b.(string); ok {
			return a.(string) + b, true
		}
	}
	return nil, false
}

func addWithReflection(a, b interface{}) (interface{}, bool) {
	if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
		return a.(int) + b.(int), true
	} else if reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Float64 {
		return a.(float64) + b.(float64), true
	} else if reflect.TypeOf(a).Kind() == reflect.String && reflect.TypeOf(b).Kind() == reflect.String {
		return a.(string) + b.(string), true
	}
	return nil, false
}

func addUsingGenerics[T int | float64 | string](a, b T) T {
	return a + b
}

type possibleType interface {
	int | float64 | string
}

func addUsingGenericsWithTypeSet[T possibleType](a, b T) T {
	return a + b
}
