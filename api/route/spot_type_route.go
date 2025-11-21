package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewSpotTypeRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.SpotTypeController{
		SpotTypeRepository: &usecase.SpotTypeUseCase{},
	}
	SpotTypeRouter := group.Group("/spotType")
	SpotTypeRouter.POST("/", ec.Create)
	SpotTypeRouter.GET("/", ec.Fetch)
	SpotTypeRouter.GET("/:id", ec.FetchById)
	SpotTypeRouter.PUT("/", ec.Update)
	SpotTypeRouter.DELETE("/:id", ec.Delete)
}
