package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewTimeSlotRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.TimeSlotController{
		TimeSlotRepository: &usecase.TimeSlotUseCase{},
	}
	TimeSlotRouter := group.Group("/timeSlot")
	TimeSlotRouter.POST("/", ec.Create)
	TimeSlotRouter.GET("/", ec.Fetch)
	TimeSlotRouter.GET("/:id", ec.FetchById)
	TimeSlotRouter.PUT("/", ec.Update)
	TimeSlotRouter.DELETE("/:id", ec.Delete)
}
