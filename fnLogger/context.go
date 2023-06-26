package fnLogger

import (
	"context"
	"fmt"
	"github.com/d3v-friends/pure-go/vars"
)

func Get(ctx context.Context, iDefault ...IfLogger) (logger IfLogger) {
	var isOk bool
	if logger, isOk = ctx.Value(vars.CTX_LOGGER).(IfLogger); !isOk {
		if len(iDefault) == 0 {
			panic(fmt.Errorf("not found logger"))
		}
		return iDefault[0]
	}
	return
}

func Set(ctx context.Context, logger IfLogger) context.Context {
	return context.WithValue(ctx, vars.CTX_LOGGER, logger)
}
