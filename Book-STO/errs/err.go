package errs

import "net/http"

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

func BadRequestError(mess string) *AppError {

	return &AppError{

		Message: mess,
		Code:    http.StatusBadRequest,
	}
}

func ErrorReadRequestBody() *AppError {

	return &AppError{

		Message: "Cannot read request body",
		Code:    http.StatusBadGateway,
	}
}

func InternalServerError(mess string) *AppError {

	return &AppError{

		Message: mess,
		Code:    http.StatusInternalServerError,
	}
}

func ServiceUnavailableError(mess string) *AppError {

	return &AppError{

		Message: mess,
		Code:    http.StatusServiceUnavailable,
	}
}

func NotFoundError(mess string) *AppError {

	return &AppError{

		Message: mess,
		Code:    http.StatusNotFound,
	}
}

func RequestTimeoutError(mess string) *AppError {

	return &AppError{

		Message: mess,
		Code:    http.StatusRequestTimeout,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewUnauthenticatedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func IsError(errors ...*AppError) *AppError {

	for i := 0; i < len(errors); i++ {

		if errors[i] != nil {

			return errors[i]
		}
	}
	return nil
}

func HaveError(errors ...*AppError) (interface{}, *AppError) {

	for i := 0; i < len(errors); i++ {

		if errors[i] != nil {

			return nil, errors[i]
		}
	}
	return nil, nil
}

func ErrorGetData() *AppError {

	return &AppError{

		Message: "Error get data",
		Code:    http.StatusInternalServerError,
	}
}

func ErrorReadData() *AppError {

	return &AppError{

		Message: "Error read data",
		Code:    http.StatusInternalServerError,
	}
}

func ErrorDeleteData() *AppError {

	return &AppError{

		Message: "Error delete data",
		Code:    http.StatusInternalServerError,
	}
}

func ErrorInsertData() *AppError {

	return &AppError{

		Message: "Error insert data",
		Code:    http.StatusInternalServerError,
	}
}

func ErrorUpdateData() *AppError {

	return &AppError{

		Message: "Error update data",
		Code:    http.StatusInternalServerError,
	}
}

func ErrorDataNotSurvive() *AppError {

	return &AppError{

		Message: "Have no data",
		Code:    http.StatusBadRequest,
	}
}
