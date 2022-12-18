package runtransaction

import (
	"gogen-transaction/domain_belajartransaction/model/repository"
	repository2 "gogen-transaction/shared/model/repository"
)

type Outport interface {
	repository2.WithTransactionDB
	repository.SaveProductRepo
	repository.SaveOrderRepo
}
