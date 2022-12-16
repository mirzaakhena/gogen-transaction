package restapi

import (
	"gogen-transaction/shared/gogen"
	"gogen-transaction/shared/infrastructure/config"
	"gogen-transaction/shared/infrastructure/logger"
	"gogen-transaction/shared/infrastructure/token"

	"github.com/gin-gonic/gin"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log      logger.Logger
	cfg      *config.Config
	jwtToken token.JWTToken
}

func NewGinController(log logger.Logger, cfg *config.Config, tk token.JWTToken) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
		jwtToken:       tk,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	resource := router.Group("/api/v1", r.authentication())
	resource.POST("/trx", r.authorization(), r.runTransactionHandler())
}
