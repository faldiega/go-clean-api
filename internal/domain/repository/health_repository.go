package repository

import "context"

type HealthRepository interface {
	PingDatabase(ctx context.Context) error
}
