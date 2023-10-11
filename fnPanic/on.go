package fnPanic

import "github.com/d3v-friends/go-pure/fnLogger"

func On(err error) {
	if err != nil {
		var logger = fnLogger.NewDefaultLogger()
		logger.Fatal("err: %s", err.Error())
		panic(err)
	}
}

func OnPointer[T any](value *T, err error) *T {
	if err != nil {
		var logger = fnLogger.NewDefaultLogger()
		logger.Fatal("err: %s", err.Error())
		panic(err)
	}
	return value
}

func OnValue[T any](value T, err error) T {
	if err != nil {
		var logger = fnLogger.NewDefaultLogger()
		logger.Fatal("err: %s", err.Error())
		panic(err)
	}
	return value
}

func Get[T any](value T, err error) T {
	if err != nil {
		var logger = fnLogger.NewDefaultLogger()
		logger.Fatal("err: %s", err.Error())
	}
	return value
}

func IsTrue(v bool, err error) {
	if !v {
		var logger = fnLogger.NewDefaultLogger()
		logger.Fatal("err: %s", err.Error())
	}
}
