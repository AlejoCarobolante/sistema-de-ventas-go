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
	NewVehicleRouter(env, timeout, publicRouter)
	NewParkingRouter(env, timeout, publicRouter)
	NewSpotRouter(env, timeout, publicRouter)
	NewPenaltyRouter(env, timeout, publicRouter)
	NewPaymentRouter(env, timeout, publicRouter)
	NewRateRouter(env, timeout, publicRouter)
	NewTimeSlotRouter(env, timeout, publicRouter)
	NewSpotTypeRouter(env, timeout, publicRouter)
	NewVehicleTypeRouter(env, timeout, publicRouter)
	NewReservationStatusRouter(env, timeout, publicRouter)
	NewReservationTypeRouter(env, timeout, publicRouter)
	NewCompatibilityRuleRouter(env, timeout, publicRouter)
}
