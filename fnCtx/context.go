package fnCtx

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-pure/fnPanic"
)

func Get[DATA any](ctx context.Context, key string) (res DATA, err error) {
	var isOk bool
	if res, isOk = ctx.Value(key).(DATA); !isOk {
		err = fmt.Errorf("not found value in context: key=%s", key)
		return
	}
	return
}

func GetP[DATA any](ctx context.Context, key string) (res DATA) {
	res = fnPanic.Get(Get[DATA](ctx, key))
	return
}

func Set[DATA any](ctx context.Context, key string, v DATA) context.Context {
	return context.WithValue(ctx, key, v)
}
