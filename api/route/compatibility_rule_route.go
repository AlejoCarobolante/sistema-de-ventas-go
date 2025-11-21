package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewCompatibilityRuleRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.CompatibilityRuleController{
		CompatibilityRuleRepository: &usecase.CompatibilityRuleUseCase{},
	}
	CompatibilityRuleRouter := group.Group("/compatibilityRule")
	CompatibilityRuleRouter.POST("/", ec.Create)
	CompatibilityRuleRouter.GET("/", ec.Fetch)
	CompatibilityRuleRouter.GET("/:id", ec.FetchById)
	CompatibilityRuleRouter.PUT("/", ec.Update)
	CompatibilityRuleRouter.DELETE("/:id", ec.Delete)
}
