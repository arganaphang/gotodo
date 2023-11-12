package todo

import (
	"application/domain"
	"context"

	"github.com/google/uuid"
)

type service struct {
	repo domain.TodoRepository
}

func NewService(repo domain.TodoRepository) domain.TodoService {
	return &service{repo: repo}
}

func (s *service) CreateTodo(ctx context.Context, todo domain.Todo) (*domain.Todo, error) {
	return s.repo.CreateTodo(ctx, todo)
}

func (s *service) GetTodos(ctx context.Context) ([]domain.Todo, error) {
	return s.repo.GetTodos(ctx)
}

func (s *service) GetTodoByID(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
	return s.repo.GetTodoByID(ctx, id)
}

func (s *service) UpdateTodoByID(ctx context.Context, todo domain.Todo) (*domain.Todo, error) {
	return s.repo.UpdateTodoByID(ctx, todo)
}

func (s *service) DeleteTodoByID(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteTodoByID(ctx, id)
}
