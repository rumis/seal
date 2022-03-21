package options

import (
	"context"
	"fmt"
	"time"
)

// ExecLogFunc is called each time when a SQL statement is executed.
// The "ts" parameter gives the time that the SQL statement takes to execute,
// sql and args refer the exec params
// err refer to the result of the execution.
type ExecLogFunc func(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error)

// ConsoleExecLogFunc print the message to console
func ConsoleExecLogFunc(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error) {
	fmt.Printf("exec(query) log, sql:%s \n args:%+v \n timespan:%dns \n error:%+v\n", sql, args, ts.Nanoseconds(), err)
}

// BuildLogFunc is called each time when we build a sql statement and args.
// The "ts" parameter gives the time that the build takes,
// while sql,args and err refer to the result of the execution.
type BuildLogFunc func(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error)

func ConsoleBuildLogFunc(ctx context.Context, ts time.Duration, sql string, args []interface{}, err error) {
	fmt.Printf("sql build log, sql:%s \n args:%+v \n timespan:%dns \n error:%+v\n", sql, args, ts.Nanoseconds(), err)
}
