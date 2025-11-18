package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewClientRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ClientController{
		ClientRepository: &usecase.ClientUseCase{},
	}
	ClientRouter := group.Group("/client")
	ClientRouter.POST("/", ec.Create)
	ClientRouter.GET("/", ec.Fetch)
	ClientRouter.GET("/:id", ec.FetchById)
	ClientRouter.PUT("/", ec.Update)
	ClientRouter.DELETE("/:id", ec.Delete)
}
