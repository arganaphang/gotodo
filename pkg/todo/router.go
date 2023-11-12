package todo

import (
	"application/domain"
	"application/helper/pagination"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	// ? Serialize
	body := domain.Todo{}
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to serialize request body",
		})
		return
	}

	// ? Validate
	// ? Validate Empty
	if len(strings.Trim(body.Title, " ")) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "title can't be empty",
		})
		return
	}

	// ? Do~
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
	paginate, err := pagination.Transform(ctx.Request.URL.Query())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "pagination invalid",
		})
		return
	}

	todos, count, err := r.service.GetTodos(ctx, paginate.Limit, paginate.Offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get todo list",
		})
		return
	}
	paginate.Finish(count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get todo list success",
		"data":    todos,
		"meta":    &paginate,
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
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "todo not found",
			})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "failed to get todo by id",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
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

	// ? Validate
	// ? Validate Empty
	if len(strings.Trim(body.Title, " ")) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "title can't be empty",
		})
		return
	}

	body.ID = id

	todo, err := r.service.UpdateTodoByID(ctx, body)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "todo not found",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update todo",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})
}
