package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewVehicleTypeRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.VehicleTypeController{
		VehicleTypeRepository: &usecase.VehicleTypeUseCase{},
	}
	VehicleTypeRouter := group.Group("/vehicleType")
	VehicleTypeRouter.POST("/", ec.Create)
	VehicleTypeRouter.GET("/", ec.Fetch)
	VehicleTypeRouter.GET("/:id", ec.FetchById)
	VehicleTypeRouter.PUT("/", ec.Update)
	VehicleTypeRouter.DELETE("/:id", ec.Delete)
}
