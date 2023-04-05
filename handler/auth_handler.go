package handler

import (
	"net/http"
	"time"
	"todolist/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const sessionKey = "session_id"

// 用户登陆
func Login(sessionStorage *model.SessionStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")
		password := c.Query("password")

		// 校验是否合法用户
		v, ok := model.UserMap[user]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			return
		}
		if password != v.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名密码错误"})
			return
		}

		// 合法用户，成功登陆，往客户端写cookie并创建一条对应session
		sessionId := uuid.New().String()
		c.SetCookie(sessionKey, sessionId, int(time.Hour.Seconds()), "/", "localhost", false, true)
		sessionStorage.Put(sessionId, v)
		c.JSON(http.StatusOK, gin.H{"message": "成功登陆"})
	}
}

// 用户登出
func Logout(sessionStorage *model.SessionStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie(sessionKey)
		if err != nil {
			// http request中不存在这个cookie，用户未登陆
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "未登陆用户不能登出"})
			return
		}

		sessionStorage.Del(sessionId)
		// 通过把MaxAge设置为-1来删除客户端cookie
		c.SetCookie(sessionKey, sessionId, -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "成功登出"})
	}
}
