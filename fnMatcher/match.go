package fnMatcher

type (
	FnMatcher[T any] func(v T) bool
)

func Has[T any](ls []T, matcher FnMatcher[T]) bool {
	for _, v := range ls {
		if matcher(v) {
			return true
		}
	}
	return false
}
