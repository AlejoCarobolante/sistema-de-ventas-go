package route

import (
	"gorm-template/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	//Todas las API Publicas
	publicRouter := gin.Group("/api")

	//Middleware para verificar AccessToken
	//protectedRouter := gin.Group("/api")

	//Todas las API Privadas
	NewReservationRouter(env, timeout, publicRouter)
	NewClientRouter(env, timeout, publicRouter)
}
