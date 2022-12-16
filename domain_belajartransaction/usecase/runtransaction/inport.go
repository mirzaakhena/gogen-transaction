package runtransaction

import (
	"gogen-transaction/domain_belajartransaction/model/entity"
	"gogen-transaction/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.ProductCreateRequest
	entity.OrderCreateRequest
}

type InportResponse struct {
}
