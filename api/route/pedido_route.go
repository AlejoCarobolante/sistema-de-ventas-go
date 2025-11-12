package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewPedidoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.PedidoController{
		PedidoRepository: &usecase.PedidoUseCase{},
	}
	PedidoRouter := group.Group("/pedido")
	PedidoRouter.POST("/", ec.Create)
	PedidoRouter.GET("/", ec.Fetch)
	PedidoRouter.GET("/:id", ec.FetchById)
	PedidoRouter.PUT("/", ec.Update)
	PedidoRouter.DELETE("/:id", ec.Delete)
}
