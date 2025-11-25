package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewPaymentRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.PaymentController{
		PaymentRepository: &usecase.PaymentUseCase{},
	}
	PaymentRouter := group.Group("/payment")
	PaymentRouter.POST("/", ec.Create)
	PaymentRouter.GET("/", ec.Fetch)
	PaymentRouter.GET("/:id", ec.FetchById)
	PaymentRouter.PUT("/", ec.Update)
	PaymentRouter.DELETE("/:id", ec.Delete)
}
