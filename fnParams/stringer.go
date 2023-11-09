package fnParams

import (
	"fmt"
	"reflect"
)

func ToString(v any) (str string, err error) {
	switch i := v.(type) {
	case fmt.Stringer:
		str = i.String()
	case fmt.GoStringer:
		str = i.GoString()
	case string:
		str = i
	default:
		err = fmt.Errorf("invalid str type: name=%s", reflect.TypeOf(v).Name())
	}
	return
}
