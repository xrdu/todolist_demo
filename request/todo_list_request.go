package request

import "time"

// TodoListCreateRequest 创建TodoList的请求体
type TodoListCreateRequest struct {
	Title      string    `json:"title" binding:"required"`
	Body       string    `json:"body" binding:"required"`
	CreateUser string    `json:"create_user"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// TodoListUpdateRequest 更新TodoList的请求体
type TodoListUpdateRequest struct {
	Id    string `json:"id" binding:"required"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
