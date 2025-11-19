package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewVehicleRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.VehicleController{
		VehicleRepository: &usecase.VehicleUseCase{},
	}
	VehicleRouter := group.Group("/vehicle")
	VehicleRouter.POST("/", ec.Create)
	VehicleRouter.GET("/", ec.Fetch)
	VehicleRouter.GET("/:id", ec.FetchById)
	VehicleRouter.PUT("/", ec.Update)
	VehicleRouter.DELETE("/:id", ec.Delete)
}
