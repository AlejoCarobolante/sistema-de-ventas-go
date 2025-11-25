package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewClientRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	cc := &controller.ClientController{
		ClientRepository: &usecase.ClientUseCase{},
	}
	ClientRouter := group.Group("/client")
	ClientRouter.POST("", cc.Create)
	ClientRouter.GET("", cc.Fetch)
	ClientRouter.GET(":id", cc.FetchByID)
	ClientRouter.PUT("", cc.Update)
	ClientRouter.DELETE(":id", cc.Delete)
}
