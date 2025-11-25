package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewParkingRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ParkingController{
		ParkingRepository: &usecase.ParkingUseCase{},
	}
	ParkingRouter := group.Group("/parking")
	ParkingRouter.POST("/", ec.Create)
	ParkingRouter.GET("/", ec.Fetch)
	ParkingRouter.GET("/:id", ec.FetchById)
	ParkingRouter.PUT("/", ec.Update)
	ParkingRouter.DELETE("/:id", ec.Delete)
}
