package logx

// Adapter interface for logger interoperability
type Adapter interface {
	Logger
	WithFormatter(formatter Formatter) Adapter
}

// Formatter interface for log output formatting
type Formatter interface {
	Format(message *string, fields map[string]any) any
}

// Log interface for log compliance (https://pkg.go.dev/log)
type Log interface {
	// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
	Fatal(v ...any)
	// Panic is equivalent to l.Print() followed by a call to panic().
	Panic(v ...any)
	// Print to the logx. Arguments are handled in the manner of fmt.Print.
	Print(v ...any)
}

// LogF interface for formatted log compliance (https://pkg.go.dev/log)
type LogF interface {
	// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
	Fatalf(format string, v ...any)
	// Panicf is equivalent to l.Printf() followed by a call to panic().
	Panicf(format string, v ...any)
	// Printf to the logx. Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...any)
}

// Logger abstraction for structured/field logging
type Logger interface {
	Log
	LogF
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Notice(v ...any)
	Noticef(format string, v ...any)
	Warning(v ...any)
	Warningf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}

// log level values
const (
	LogLevelDebug   int = 100
	LogLevelInfo    int = 200
	LogLevelNotice  int = 300
	LogLevelWarning int = 400
	LogLevelError   int = 500
	LogLevelFatal   int = 600
	LogLevelPanic   int = 700
)

// special attribute names
const (
	FieldNameLogLevel     string = "level"
	FieldNameMessage      string = "message"
	FieldNameTimestamp    string = "timestamp"
	FieldNameCallerFunc   string = "caller:func"
	FieldNameCallerFile   string = "caller:file"
	FieldNameCallerLine   string = "caller:line"
	FieldNameHttpRequest  string = "http:request"
	FieldNameHttpResponse string = "http:response"
)
