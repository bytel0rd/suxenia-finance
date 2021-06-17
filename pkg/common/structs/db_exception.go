package structs

type DBException struct {
	Exception         error
	isValidationError bool
	IsExecError       bool
}

func (e *DBException) Error() string {
	return e.Exception.Error()
}

func NewDBException(exception error, isValidationError bool) DBException {

	return DBException{
		Exception:         exception,
		isValidationError: isValidationError,
		IsExecError:       !isValidationError,
	}

}
