package options

import (
	"context"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

// ExecLogFunc is called each time when a SQL statement is executed.
// The "ts" parameter gives the time that the SQL statement takes to execute,
// sql and args refer the exec params
// err refer to the result of the execution.
type ExecLogFunc func(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error)

// ConsoleExecLogFunc print the message to console
func ConsoleExecLogFunc(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error) {
	traceId := ctx.Value(DefaultTraceKey)
	fmt.Printf("exec(query) log: \n trace:%v\n  sql:%s \n  args:%+v \n  timespan:%dns \n  error:%+v\n \n", traceId, sql, args, ts.Nanoseconds(), err)
}

// BuildLogFunc is called each time when we build a sql statement and args.
// The "ts" parameter gives the time that the build takes,
// while sql,args and err refer to the result of the execution.
type BuildLogFunc func(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error)

// ConsoleBuildLogFunc 控制台日志输出
func ConsoleBuildLogFunc(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error) {
	traceId := ctx.Value(DefaultTraceKey)
	fmt.Printf("sql build log: \n  trace:%v \n  sql:%s \n  args:%+v \n  timespan:%dns \n  error:%+v\n \n", traceId, sql, args, ts.Nanoseconds(), err)
}

// 链路追踪KEY
type TraceKey struct{}

// 默认的TraceKey
var DefaultTraceKey TraceKey = TraceKey{}

// NewTraceId 生成新的UUID
func NewTraceId() string {
	u := uuid.NewV4()
	return strings.Replace(u.String(), "-", "", -1)
}
