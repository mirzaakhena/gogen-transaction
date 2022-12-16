package runtransaction

import (
	"gogen-transaction/domain_belajartransaction/model/repository"
)

type Outport interface {
	repository.SaveProductRepo
	repository.SaveOrderRepo
}
