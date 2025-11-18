package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewReservationRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ReservationController{
		ReservationRepository: &usecase.ReservationUseCase{},
	}
	ReservationRouter := group.Group("/reservation")
	ReservationRouter.POST("/", ec.Create)
	ReservationRouter.GET("/", ec.Fetch)
	ReservationRouter.GET("/:id", ec.FetchById)
	ReservationRouter.PUT("/", ec.Update)
	ReservationRouter.DELETE("/:id", ec.Delete)
}
