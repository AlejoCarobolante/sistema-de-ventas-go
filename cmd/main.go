package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm-template/api/route"
	"gorm-template/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()
	router.Use(cors.Default())

	route.Setup(env, timeout, router)

	if err := router.Run(env.ServerAddress); err != nil {
		log.Fatalf("error al iniciar el servidor: %v", err)
	}
}
