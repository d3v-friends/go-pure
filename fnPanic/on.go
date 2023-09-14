package fnPanic

func On(err error) {
	if err != nil {
		panic(err)
	}
}

func OnPointer[T any](value *T, err error) *T {
	if err != nil {
		panic(err)
	}
	return value
}

func OnValue[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Get[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func IsTrue(v bool, err error) {
	if !v {
		panic(err)
	}
}
