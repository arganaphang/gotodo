package main

import (
	"application/pkg/todo"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
)

var PORT string

func init() {
	logger, err := zap.NewDevelopment()
	if os.Getenv("MODE") == "PRODUCTION" {
		logger, err = zap.NewProduction()
	}
	zap.ReplaceGlobals(zap.Must(logger, err))

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
}

func main() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgres://postgres:postgres@localhost:5432/todos?sslmode=disable")))
	db := bun.NewDB(sqldb, pgdialect.New())

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	// ? Health Check Endpoint
	app.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	todoRepo := todo.NewRepository(db)
	todoSvc := todo.NewService(todoRepo)
	todo.NewRouter(app, todoSvc)

	app.Run(fmt.Sprintf("0.0.0.0:%s", PORT))
}
