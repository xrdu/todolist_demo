package model

type User struct {
	Name     string
	Password string
	Role     string
}

// UserMap 合法用户全集
var UserMap = map[string]*User{
	"tom": {
		Name:     "tom",
		Password: "tom123",
		Role:     "admin",
	},
	"mary": {
		Name:     "mary",
		Password: "mary123",
		Role:     "user",
	},
	"jack": {
		Name:     "jack",
		Password: "jack123",
		Role:     "user",
	},
}
