package todo

import (
	"application/domain"
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) domain.TodoRepository {
	return &repository{db: db}
}

func (r *repository) CreateTodo(ctx context.Context, todo domain.Todo) (*domain.Todo, error) {
	if _, err := r.db.NewInsert().Model(&todo).Column("title").Returning("*").Exec(ctx); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *repository) GetTodos(ctx context.Context, limit, offset int) ([]domain.Todo, int, error) {
	todos := []domain.Todo{}
	count, err := r.db.NewSelect().Model(&todos).Limit(limit).Offset(offset).Order("created_at desc").ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}
	return todos, count, nil
}

func (r *repository) GetTodoByID(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
	todo := &domain.Todo{}
	if err := r.db.NewSelect().Model(todo).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *repository) UpdateTodoByID(ctx context.Context, todo domain.Todo) (*domain.Todo, error) {
	if _, err := r.GetTodoByID(ctx, todo.ID); err != nil {
		return nil, err
	}

	if _, err := r.db.NewUpdate().Model(&todo).Column("title", "is_done").Where("id = ?", todo.ID).Returning("*").Exec(ctx); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *repository) DeleteTodoByID(ctx context.Context, id uuid.UUID) error {
	if _, err := r.GetTodoByID(ctx, id); err != nil {
		return err
	}

	_, err := r.db.NewDelete().Model((*domain.Todo)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
