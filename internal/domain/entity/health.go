package entity

import "time"

type HealthStatus string

const (
	HealthUp   HealthStatus = "up"
	HealthDown HealthStatus = "down"
)

type ComponentHealth struct {
	Status    HealthStatus
	LatencyMs int64
}

type Health struct {
	Status     HealthStatus
	Timestamp  time.Time
	Components map[string]ComponentHealth
}
