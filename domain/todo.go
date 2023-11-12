package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos"`
	ID            uuid.UUID `bun:"id,pk" json:"id"`
	Title         string    `bun:"title" json:"title"`
	IsDone        bool      `bun:"is_done" json:"is_done"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
}

type TodoService interface {
	CreateTodo(ctx context.Context, todo Todo) (*Todo, error)
	GetTodos(ctx context.Context) ([]Todo, error)
	GetTodoByID(ctx context.Context, id uuid.UUID) (*Todo, error)
	UpdateTodoByID(ctx context.Context, todo Todo) (*Todo, error)
	DeleteTodoByID(ctx context.Context, id uuid.UUID) error
}

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo Todo) (*Todo, error)
	GetTodos(ctx context.Context) ([]Todo, error)
	GetTodoByID(ctx context.Context, id uuid.UUID) (*Todo, error)
	UpdateTodoByID(ctx context.Context, todo Todo) (*Todo, error)
	DeleteTodoByID(ctx context.Context, id uuid.UUID) error
}
