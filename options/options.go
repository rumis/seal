package seal

import (
	"context"
	"reflect"
	"time"
)

type EncodeHookFunc func(typ reflect.Type, data interface{}) (interface{}, error)

type ExecLogFunc func(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error)

type SealOptions struct {
	EncodeHook EncodeHookFunc
}
