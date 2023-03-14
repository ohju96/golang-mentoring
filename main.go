package main

import (
	"boardgame/app/config/db"
	toml "boardgame/app/config/toml"
	"boardgame/app/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	app := gin.Default()

	config := toml.InitToml("./app/config/toml/config.toml")

	Init(config, app)

	app.Run()

}

func Init(config toml.Config, app *gin.Engine) {
	db.InitMySQL(&config)
	router.InitRouter(&config, app)
}
