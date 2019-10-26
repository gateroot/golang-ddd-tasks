package model

import "context"

type CreateRequest struct {
	Title string
}

type Service interface {
	CreateTask(ctx context.Context, r CreateRequest) error
}

type service struct {
	r Repository
}

func (s service) CreateTask(ctx context.Context, r CreateRequest) error {
	t := NewTask(r.Title)
}

