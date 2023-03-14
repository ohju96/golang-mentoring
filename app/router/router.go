package router

import (
	"boardgame/app/config/toml"
	"github.com/gin-gonic/gin"
)

func InitRouter(c *toml.Config, g *gin.Engine) {
	InitUserRouter(c, g)
}
