package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewReservationTypeRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ReservationTypeController{
		ReservationTypeRepository: &usecase.ReservationTypeUseCase{},
	}
	ReservationTypeRouter := group.Group("/reservationType")
	ReservationTypeRouter.POST("/", ec.Create)
	ReservationTypeRouter.GET("/", ec.Fetch)
	ReservationTypeRouter.GET("/:id", ec.FetchById)
	ReservationTypeRouter.PUT("/", ec.Update)
	ReservationTypeRouter.DELETE("/:id", ec.Delete)
}
