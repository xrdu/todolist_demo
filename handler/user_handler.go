package handler

import (
	"net/http"
	"todolist/model"
	"todolist/request"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// CreateOrUpdateUser godoc
// @Summary      创建或更新User
// @Description  创建或更新User
// @Tags         createuser
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /user [post]
func (h *UserHandler) CreateOrUpdateUser(c *gin.Context) {
	var req request.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "req body is invalid"})
		return
	}
	var user model.User
	user.Name = req.Name
	user.Password = req.Password
	user.Role = req.Role
	model.UserMap[user.Name] = &user
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// DeleteUser godoc
// @Summary      删除User
// @Description  删除User
// @Tags         createuser
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /user [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	username := c.Query("name")
	if username != "" {
		delete(model.UserMap, username)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    "deleted user " + username,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "用户不存在"})
	}
}
