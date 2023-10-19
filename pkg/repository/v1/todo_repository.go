package repository

import (
	"context"
	"errors"
	"grpc_bri/pkg/infrastructure/v1"
	"sync"
)

var (
	instance ITodoRepository
	once     sync.Once
)

type ITodoRepository interface {
	CreateTodo(ctx context.Context, todo *TodoModel) error
	ReadTodo(ctx context.Context, id string) (*TodoModel, error)
	ReadAll(ctx context.Context) (*[]TodoModel, error)
	UpdateTodo(ctx context.Context, todo *TodoModel) error
	Delete(ctx context.Context, id string) error
}

func New() ITodoRepository {
	once.Do(func() {
		instance = &TodoRepository{}
	})
	return instance
}

type TodoRepository struct {
}

func (t *TodoRepository) CreateTodo(ctx context.Context, todo *TodoModel) error {
	if err := infrastructure.Connect().WithContext(ctx).Create(todo).Error; err != nil {
		return errors.New("failed to create")
	}
	return nil
}

func (t *TodoRepository) ReadTodo(ctx context.Context, id string) (*TodoModel, error) {
	var todo TodoModel
	if err := infrastructure.Connect().WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, errors.New("failed to read")
	}
	return &todo, nil
}

func (t *TodoRepository) ReadAll(ctx context.Context) (*[]TodoModel, error) {
	var todos []TodoModel
	if err := infrastructure.Connect().WithContext(ctx).Find(&todos).Error; err != nil {
		return nil, errors.New("failed to read all")
	}
	return &todos, nil
}

func (t *TodoRepository) UpdateTodo(ctx context.Context, todo *TodoModel) error {
	result := infrastructure.Connect().WithContext(ctx).Updates(todo)
	if result.Error != nil || result.RowsAffected < 1 {
		return errors.New("failed to update")
	}
	return nil
}

func (t *TodoRepository) Delete(ctx context.Context, id string) error {
	result := infrastructure.Connect().WithContext(ctx).Delete(&TodoModel{}, id)
	if result.Error != nil || result.RowsAffected < 1 {
		return errors.New("failed to delete")
	}
	return nil
}
