package errorenum

import "gogen-transaction/shared/model/apperror"

const (
	SomethingError   apperror.ErrorType = "ER0000 something error"
	UserMustNotEmpty apperror.ErrorType = "ER0001 user must not empty"
)
