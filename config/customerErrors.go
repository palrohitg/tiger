package config

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func SendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	response := ErrorResponse{
		Status:  "error",
		Message: message,
	}
	ctx.JSON(statusCode, response)
}

func SendSuccessResponse(ctx *gin.Context, statusCode int, message string, data map[string]interface{}) {
	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	ctx.JSON(statusCode, response)
}

type CustomError struct {
	Code    int    `json:"-"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (c CustomError) Error() string {
	return c.Message
}
func ValidationErrorReal(message string) CustomError {
	if message == "" {
		message = ValidationErrMsg
	}

	error_ := createError(message, 400)
	return error_
}
func ValidationError(message string) CustomError {
	if message == "" {
		message = ValidationErrMsg
	}

	error_ := createError(message, 400)
	return error_
}

func createError(message string, code int) CustomError {
	error_ := CustomError{Code: code, Message: message, Status: ErrStatus}
	return error_
}

func PermissionDenied(message string) CustomError {
	if message == "" {
		message = PermissionDeniedErrMsg
	}
	error_ := createError(message, 403)
	return error_
}
