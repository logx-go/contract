package logx

// Adapter interface for logger interoperability
type Adapter interface {
	Logger
	// WithFormatter returned a clone of the adapter and sets given formatter to it
	WithFormatter(formatter Formatter) Adapter
}

// Formatter interface for log output formatting
type Formatter interface {
	// Format formats log output before it's passed to the wrapped logger
	// it returns the formatted message as string and formatted fields untyped
	Format(message *string, fields map[string]any) (messageF string, fieldsF map[string]any)
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
	// Debug sets the log level to LogLevelDebug and log given values to Print()
	Debug(v ...any)
	// Debugf sets the log level to LogLevelDebug and log given values to Printf()
	Debugf(format string, v ...any)
	// Info sets the log level to LogLevelInfo and log given values to Print()
	Info(v ...any)
	// Infof sets the log level to LogLevelInfo and log given values to Printf()
	Infof(format string, v ...any)
	// Notice sets the log level to LogLevelNotice and log given values to Print()
	Notice(v ...any)
	// Noticef sets the log level to LogLevelNotice and log given values to Printf()
	Noticef(format string, v ...any)
	// Warning sets the log level to LogLevelWarning and log given values to Print()
	Warning(v ...any)
	// Warningf sets the log level to LogLevelWarning and log given values to Printf()
	Warningf(format string, v ...any)
	// Error sets the log level to LogLevelError and log given values to Print()
	Error(v ...any)
	// Errorf sets the log level to LogLevelError and log given values to Printf()
	Errorf(format string, v ...any)
	// WithField clones the current logger and adds a new field to it
	WithField(name string, value any) Logger
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
	FieldNameHTTPRequest  string = "http:request"
	FieldNameHTTPResponse string = "http:response"
)
