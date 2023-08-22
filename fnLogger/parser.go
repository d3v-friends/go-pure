package fnLogger

import "log"

type (
	DefaultPrinter struct{}
)

func (x *DefaultPrinter) Print(v Fields) {
	log.Print(v.ToJson())
}

func NewDefaultLogger() IfLogger {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return NewLogger(&DefaultPrinter{})
}
