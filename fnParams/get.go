package fnParams

func Get[T any](v []T) T {
	if len(v) == 0 {
		return *new(T)
	}
	return v[0]
}
