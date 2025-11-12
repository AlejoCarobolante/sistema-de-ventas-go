package main

import (
	"gorm-template/bootstrap"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	route "gorm-template/api/route"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	gin.Use(cors.Default())

	route.Setup(env, timeout, gin)
	gin.Run(env.ServerAddress)
}
