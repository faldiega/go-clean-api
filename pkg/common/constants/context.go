package constants

// pakai string biasa → rawan konflik
// jadi pakai custom type → tidak akan konflik
type contextKey string

const (
	TraceIDKey contextKey = "trace_id"
	LoggerKey  contextKey = "logger"
)
