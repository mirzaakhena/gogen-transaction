package restapi

import (
	"context"
	"gogen-transaction/domain_belajartransaction/usecase/runtransaction"
	"gogen-transaction/shared/gogen"
	"gogen-transaction/shared/infrastructure/logger"
	"gogen-transaction/shared/model/payload"
	"gogen-transaction/shared/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *ginController) runTransactionHandler() gin.HandlerFunc {

	type InportRequest = runtransaction.InportRequest
	type InportResponse = runtransaction.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		//var jsonReq request
		//err := c.BindJSON(&jsonReq)
		//if err != nil {
		//	r.log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		//	return
		//}

		id := util.GenerateID(5)
		now := time.Now()

		var req InportRequest
		req.ProductCreateRequest.Now = now
		req.ProductCreateRequest.RandomString = id
		req.OrderCreateRequest.Now = now
		req.OrderCreateRequest.RandomString = id
		req.User = c.DefaultQuery("user", "")

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
