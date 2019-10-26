package model

import "context"

type Repository interface {
	Create(ctx context.Context, t *Task) error
	Update(ctx context.Context, t *Task) error
	GetByID(ctx context.Context, id ID) (*Task, error)
	Delete(ctx context.Context, id ID) error
	AllocateIDs(ctx context.Context, num int) ([]string, error)
}
