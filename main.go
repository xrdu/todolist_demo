package main

import (
	"todolist/handler"
	"todolist/middleware"
	"todolist/model"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Todolist API
// @version         1.0
// @description     This is a sample
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
func main() {

	todolist := model.NewTodoList()

	todolisthandler := handler.NewTodoListHandler(todolist)

	userhandler := handler.NewUserHandler()

	sessionStorage := model.NewSessionStorage()

	r := gin.Default()
	// 业务逻辑路由
	r1 := r.Group("/")
	r1.Use(middleware.UserAuth(sessionStorage))
	r1.GET("/todo", todolisthandler.GetTodoLists)
	r1.POST("/todo", todolisthandler.CreateTodoLists)
	r1.DELETE("/todo/:id", todolisthandler.DeleteTodoLists)
	r1.POST("/updatetodo", todolisthandler.UpdateTodoLists)

	// 用户管理路由
	r2 := r.Group("/user")
	r2.Use(middleware.AdminAuth(sessionStorage))
	r2.POST("/", userhandler.CreateOrUpdateUser)
	r2.DELETE("/", userhandler.DeleteUser)

	r.GET("/login", handler.Login(sessionStorage))   // 登陆
	r.GET("/logout", handler.Logout(sessionStorage)) // 登出

	// Swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
