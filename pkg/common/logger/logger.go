package logger

import (
	"context"
	"go-simple-api/pkg/common/constants"

	"go.uber.org/zap"
)

// buat logger baru dengan trace_id
func WithTraceID(base *zap.Logger, traceID string) *zap.Logger {
	return base.With(zap.String("trace_id", traceID))
}

// simpan logger ke context
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, constants.LoggerKey, logger)
}

// ambil logger dari context
// jika tidak ada → return logger tanpa trace_id (fallback)
func FromContext(ctx context.Context, fallback *zap.Logger) *zap.Logger {
	if logger, ok := ctx.Value(constants.LoggerKey).(*zap.Logger); ok {
		return logger
	}
	return fallback
}
