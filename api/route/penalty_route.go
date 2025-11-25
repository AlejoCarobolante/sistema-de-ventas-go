package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewPenaltyRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.PenaltyController{
		PenaltyRepository: &usecase.PenaltyUseCase{},
	}
	PenaltyRouter := group.Group("/penalty")
	PenaltyRouter.POST("/", ec.Create)
	PenaltyRouter.GET("/", ec.Fetch)
	PenaltyRouter.GET("/:id", ec.FetchById)
	PenaltyRouter.PUT("/", ec.Update)
	PenaltyRouter.DELETE("/:id", ec.Delete)
}
