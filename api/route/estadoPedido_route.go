package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewEstadoPedidoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.EstadoPedidoController{
		EstadoPedidoRepository: &usecase.EstadoPedidoUseCase{},
	}
	EstadoPedidoRouter := group.Group("/estadoPedido")
	EstadoPedidoRouter.POST("/", ec.Create)
	EstadoPedidoRouter.GET("/", ec.Fetch)
	EstadoPedidoRouter.GET("/:id", ec.FetchById)
	EstadoPedidoRouter.PUT("/", ec.Update)
	EstadoPedidoRouter.DELETE("/:id", ec.Delete)
}
