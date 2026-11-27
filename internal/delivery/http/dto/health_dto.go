package dto

import (
	"time"

	"go-clean-api/internal/domain/entity"
)

type ComponentHealthResponse struct {
	Status    string `json:"status"`
	LatencyMs int64  `json:"latency_ms"`
}

type HealthResponse struct {
	Status     string                             `json:"status"`
	Timestamp  time.Time                          `json:"timestamp"`
	Components map[string]ComponentHealthResponse `json:"components,omitempty"`
}

func ToHealthResponse(h entity.Health) HealthResponse {
	res := HealthResponse{
		Status:    string(h.Status),
		Timestamp: h.Timestamp,
	}

	if len(h.Components) > 0 {
		res.Components = make(map[string]ComponentHealthResponse, len(h.Components))
		for name, c := range h.Components {
			res.Components[name] = ComponentHealthResponse{
				Status:    string(c.Status),
				LatencyMs: c.LatencyMs,
			}
		}
	}
	return res
}
