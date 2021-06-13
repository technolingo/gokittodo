package service

import (
	"context"
	"time"
)

// TodosService describes the service.
type TodosService interface {
	Create(ctx context.Context, content string, dueTime time.Time) error
	Retrieve(ctx context.Context, dueTime time.Time) error
	Update(ctx context.Context, ID uint64, content string, dueTime time.Time, complete bool) error
	Delete(ctx context.Context, ID uint64) error
}
