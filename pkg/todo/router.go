package todo

import (
	"application/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type router struct {
	service domain.TodoService
}

func NewRouter(app *gin.Engine, service domain.TodoService) {
	// ? Create Method
	r := router{service: service}
	// ? Create Route Group
	group := app.Group("/todos")
	group.POST("/", r.CreateTodo)
	group.GET("/", r.GetTodos)
	group.GET("/:id", r.GetTodoByID)
	group.PUT("/:id", r.UpdateTodoByID)
	group.DELETE("/:id", r.DeleteTodoByID)
}

func (r *router) CreateTodo(ctx *gin.Context) {
	body := domain.Todo{}
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to serialize request body",
		})
	}
	todo, err := r.service.CreateTodo(ctx, body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create todo",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "todo created",
		"data":    *todo,
	})
}

func (r *router) GetTodos(ctx *gin.Context) {
	todos, err := r.service.GetTodos(ctx)
	if err != nil {
		zap.L().Info(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get todo list",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "get todo list success",
		"data":    todos,
	})
}

func (r *router) GetTodoByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id invalid",
		})
		return
	}
	todo, err := r.service.GetTodoByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "get todo by id success",
		"data":    todo,
	})
}

func (r *router) UpdateTodoByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id invalid",
		})
		return
	}

	body := domain.Todo{}
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to serialize request body",
		})
	}

	body.ID = id

	todo, err := r.service.UpdateTodoByID(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update todo",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "todo updated",
		"data":    todo,
	})
}

func (r *router) DeleteTodoByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id invalid",
		})
		return
	}

	if err := r.service.DeleteTodoByID(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update todo",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "todo deleted",
	})
}
