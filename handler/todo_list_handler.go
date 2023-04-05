package handler

import (
	"net/http"
	"time"
	"todolist/model"
	"todolist/request"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoListHandler struct {
	todolist *model.TodoList
}

func NewTodoListHandler(todolist *model.TodoList) *TodoListHandler {
	return &TodoListHandler{
		todolist: todolist,
	}
}

// GetTodoLists godoc
// @Summary      获取所有TodoList
// @Description  获取所有TodoList
// @Tags         gettodolist
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.TodoList
// @Router       /todolists [get]
func (h *TodoListHandler) GetTodoLists(c *gin.Context) {
	// 判断是否是管理员，管理员返回全部，否则只返回该用户todolist
	user, _ := c.Get("user")
	if isAdmin(user.(*model.User)) {
		c.JSON(http.StatusOK, h.todolist.ListAll())
	} else {
		c.JSON(http.StatusOK, h.todolist.ListByUser(user.(*model.User).Name))
	}
}

// CreateTodoLists godoc
// @Summary      创建TodoList
// @Description  创建TodoList
// @Tags         createtodolist
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.TodoList
// @Router       /todo [post]
func (h *TodoListHandler) CreateTodoLists(c *gin.Context) {
	var req request.TodoListCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "req body is invalid"})
		return
	}
	var todo model.Todo
	user, _ := c.Get("user")
	todo.Id = uuid.New().String()
	todo.Title = req.Title
	todo.Body = req.Body
	todo.CreateUser = user.(*model.User).Name
	todo.CreateTime = time.Now()
	h.todolist.Create(todo)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// DeleteTodoLists godoc
// @Summary      删除TodoList
// @Description  删除TodoList
// @Tags         deletetodolist
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.TodoList
// @Router       /todo [delete]
func (h *TodoListHandler) DeleteTodoLists(c *gin.Context) {
	id := c.Param("id")
	// TODO 判断用户是否是管理员，或修改的todo是否归该用户所有，否则返回错误
	user, _ := c.Get("user")
	// 根据id获取todo
	todo := h.todolist.Get(id)
	if isAdmin(user.(*model.User)) || todo.CreateUser == user.(*model.User).Name {
		h.todolist.Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    "deleted book " + id,
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无权限"})
		return
	}

}

// UpdateTodoLists godoc
// @Summary      更新TodoList
// @Description  更新TodoList
// @Tags         updatetodolist
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.TodoList
// @Router       /updatetodo [post]
func (h *TodoListHandler) UpdateTodoLists(c *gin.Context) {
	var req request.TodoListUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "req body is invalid"})
		return
	}
	id := c.Param("id")
	todo := h.todolist.Get(id)

	// TODO 判断用户是否是管理员，或修改的todo是否归该用户所有，否则返回错误
	user, _ := c.Get("user")
	if isAdmin(user.(*model.User)) || todo.CreateUser == user.(*model.User).Name {
		if req.Title != "" {
			todo.Title = req.Title
		}
		if req.Body != "" {
			todo.Body = req.Body
		}
		todo.UpdateTime = time.Now()
		h.todolist.Update(todo)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无权限"})
		return
	}

}

func isAdmin(user *model.User) bool {
	if user.Role == "admin" {
		return true
	} else {
		return false
	}
}
