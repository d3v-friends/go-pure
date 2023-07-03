package fnReflect

import "reflect"

func ToPointer[T any](v T) *T {
	return &v
}

func IsKind(v any, kind reflect.Kind) bool {
	to := reflect.TypeOf(v)
	return to.Kind() == kind
}
