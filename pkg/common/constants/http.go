package constants

import "time"

const (
	DefaultServerWriteTimeout      = 120 * time.Second
	DefaultServerReadTimeout       = 120 * time.Second
	DefaultServerIdleTimeout       = 120 * time.Second
	DefaultServerReadHeaderTimeout = 120 * time.Second
)
