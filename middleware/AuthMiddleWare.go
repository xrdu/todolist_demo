package middleware

import (
	"net/http"
	"todolist/model"

	"github.com/gin-gonic/gin"
)

const sessionKey = "session_id"

// 用户认证中间件
func UserAuth(sessionStorage *model.SessionStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie(sessionKey)
		if err != nil {
			// http request中不存在这个cookie，用户未登陆
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未登陆"})
			return
		}

		if !sessionStorage.Exist(sessionId) {
			// cookie中的sessionId不是合法的
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "非法SessionID"})
			return
		}

		// 把对应用户的数据放在上下文中方便后续取用
		c.Set("user", sessionStorage.Get(sessionId))

		c.Next()
	}
}

func AdminAuth(sessionStorage *model.SessionStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie(sessionKey)

		if err != nil {
			// http request中不存在这个cookie，用户未登陆
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未登陆"})
			return
		}

		if !sessionStorage.Exist(sessionId) {
			// cookie中的sessionId不是合法的
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "非法SessionID"})
			return
		}

		user := sessionStorage.Get(sessionId)

		if user.Role != "admin" {
			// 没有操作权限
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无权限"})
			return
		}
		c.Next()
	}
}
