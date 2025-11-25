package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewRateRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.RateController{
		RateRepository: &usecase.RateUseCase{},
	}
	RateRouter := group.Group("/rate")
	RateRouter.POST("/", ec.Create)
	RateRouter.GET("/", ec.Fetch)
	RateRouter.GET("/:id", ec.FetchById)
	RateRouter.PUT("/", ec.Update)
	RateRouter.DELETE("/:id", ec.Delete)
}
