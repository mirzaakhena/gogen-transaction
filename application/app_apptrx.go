package application

import (
	"gogen-transaction/domain_belajartransaction/controller/restapi"
	"gogen-transaction/domain_belajartransaction/gateway/withmongodb"
	"gogen-transaction/domain_belajartransaction/usecase/runtransaction"
	"gogen-transaction/shared/gogen"
	"gogen-transaction/shared/infrastructure/config"
	"gogen-transaction/shared/infrastructure/logger"
	"gogen-transaction/shared/infrastructure/server"
	"gogen-transaction/shared/infrastructure/token"
)

type apptrx struct{}

func NewApptrx() gogen.Runner {
	return &apptrx{}
}

func (apptrx) Run() error {

	const appName = "apptrx"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withmongodb.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := restapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		runtransaction.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
