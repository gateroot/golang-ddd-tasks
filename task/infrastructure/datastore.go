package infrastructure

import (
	"cloud.google.com/go/datastore"
	"context"
	"errors"
	"github.com/gateroot/golang-ddd-tasks/task/model"
)

type Task struct {
	ID     string
	Title  string
	Status int
}

func toEntity(t model.Task) Task {
	return Task{
		ID:     t.ID().String(),
		Title:  t.Title().String(),
		Status: t.Status().ToInt(),
	}
}

func (t Task) toModel() *model.Task {
	s := model.Status(t.Status)
	m := model.NewTask(t.ID, t.Title, s)
	return m
}

type TaskRepository struct {
	client *datastore.Client
}

func (r *TaskRepository) NewClient(ctx context.Context) error {
	c, err := datastore.NewClient(ctx, datastore.DetectProjectID)
	if err != nil {
		return errors.New("error")
	}
	r.client = c
	return nil
}

func (r *TaskRepository) Create(ctx context.Context, t *model.Task) error {
	panic("implement me")
}

func (r *TaskRepository) Update(ctx context.Context, t *model.Task) error {
	panic("implement me")
}

func (r *TaskRepository) GetByID(ctx context.Context, id model.ID) (*model.Task, error) {
	panic("implement me")
}

func (r *TaskRepository) Delete(ctx context.Context, id model.ID) error {
	panic("implement me")
}

func (r *TaskRepository) AllocateIDs(ctx context.Context, num int) ([]string, error) {
	if err := r.NewClient(ctx); err != nil {
		return nil, errors.New("error")
	}
	var keys []*datastore.Key
	for i := 0; i < num; i++ {
		keys = append(keys, datastore.IncompleteKey("Task", nil))
	}
	keys, err := r.client.AllocateIDs(ctx, keys)
	if err != nil {
		return nil, errors.New("error")
	}
	
	return keys, nil
}
