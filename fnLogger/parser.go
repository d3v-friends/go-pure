package fnLogger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type (
	DefaultPrinter struct{}
)

func (x *DefaultPrinter) Print(v Fields) {
	var pc, file, line, ok = runtime.Caller(3)
	if !ok {
		file = "not_found"
		line = 0
	}

	var fnName = "not_found_fn()"
	var fn = runtime.FuncForPC(pc)
	if fn != nil {
		fnName = fmt.Sprintf("%s()", strings.TrimLeft(filepath.Ext(fn.Name()), "."))
	}

	v["code"] = fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)

	log.Print(v.ToJson())
}

func NewDefaultLogger() IfLogger {
	return NewLogger(&DefaultPrinter{})
}
