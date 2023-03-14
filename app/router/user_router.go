package router

import (
	"boardgame/app/config/toml"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(c *toml.Config, g *gin.Engine) {
	userRouter := g.Group("/api/users")
	{
		userRouter.POST("/signup") // 회원가입
	}
}
