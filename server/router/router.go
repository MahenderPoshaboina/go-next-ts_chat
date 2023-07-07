package router

import (
	"server/internal/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
}

func Start(addr string) error {
	return r.Run(addr)
}
