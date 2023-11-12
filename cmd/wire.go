//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"application/domain"
	"application/pkg/todo"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func Initialize(db *bun.DB) *domain.Services {
	wire.Build(todo.NewService, todo.NewRepository, wire.Struct(new(domain.Services), "*"))
	return nil
}
