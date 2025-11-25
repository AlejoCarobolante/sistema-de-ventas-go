package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewReservationStatusRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ReservationStatusController{
		ReservationStatusRepository: &usecase.ReservationStatusUseCase{},
	}
	ReservationStatusRouter := group.Group("/reservationStatus")
	ReservationStatusRouter.POST("/", ec.Create)
	ReservationStatusRouter.GET("/", ec.Fetch)
	ReservationStatusRouter.GET("/:id", ec.FetchById)
	ReservationStatusRouter.PUT("/", ec.Update)
	ReservationStatusRouter.DELETE("/:id", ec.Delete)
}
