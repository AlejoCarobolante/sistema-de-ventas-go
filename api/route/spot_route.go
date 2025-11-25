package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewSpotRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.SpotController{
		SpotRepository: &usecase.SpotUseCase{},
	}
	SpotRouter := group.Group("/spot")
	SpotRouter.POST("/", ec.Create)
	SpotRouter.GET("/", ec.Fetch)
	SpotRouter.GET("/:id", ec.FetchById)
	SpotRouter.PUT("/", ec.Update)
	SpotRouter.DELETE("/:id", ec.Delete)
}
