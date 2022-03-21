package options

import (
	"reflect"
	"time"
)

// EncodeHookFunc this defined a func which will be called on each value when we use a INSERT clause
type EncodeHookFunc func(typ reflect.Type, data interface{}) (interface{}, error)

// GroupEncodeHook support multi EncodeHookFunc
// The composed funcs are called in order, with the result of the
// previous transformation.
func GroupEncodeHook(fns ...EncodeHookFunc) EncodeHookFunc {
	return func(typ reflect.Type, data interface{}) (interface{}, error) {
		for _, fn := range fns {
			data, err := fn(typ, data)
			if err != nil {
				return data, err
			}
		}
		return data, nil
	}
}

// Time2StringEncodeHook converts time.Time to string which the layout
func Time2StringEncodeHook(typ reflect.Type, data interface{}) (interface{}, error) {
	if typ == reflect.TypeOf(time.Time{}) {
		tStr := data.(time.Time).Format("2006-01-02 15:04:05")
		return tStr, nil
	}
	return data, nil
}
