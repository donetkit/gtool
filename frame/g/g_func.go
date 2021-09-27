package g

import (
	"github.com/donetkit/gtool/container/gvar"
	"github.com/donetkit/gtool/internal/empty"
	"github.com/donetkit/gtool/util/gutil"
)

// NewVar returns a gvar.Var.
func NewVar(i interface{}, safe ...bool) *Var {
	return gvar.New(i, safe...)
}

// Dump dumps a variable to stdout with more manually readable.
func Dump(i ...interface{}) {
	gutil.Dump(i...)
}

// Export exports a variable to string with more manually readable.
func Export(i ...interface{}) string {
	return gutil.Export(i...)
}

// Throw throws an exception, which can be caught by TryCatch function.
func Throw(exception interface{}) {
	gutil.Throw(exception)
}

// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
func Try(try func()) (err error) {
	return gutil.Try(try)
}

// TryCatch implements try...catch... logistics using internal panic...recover.
// It automatically calls function `catch` if any exception occurs ans passes the exception as an error.
func TryCatch(try func(), catch ...func(exception error)) {
	gutil.TryCatch(try, catch...)
}

// IsNil checks whether given `value` is nil.
// Parameter <traceSource> is used for tracing to the source variable if given `value` is type
// of pinter that also points to a pointer. It returns nil if the source is nil when `traceSource`
// is true.
// Note that it might use reflect feature which affects performance a little.
func IsNil(value interface{}, traceSource ...bool) bool {
	return empty.IsNil(value, traceSource...)
}

// IsEmpty checks whether given `value` empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0.
// Or else it returns true.
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}
