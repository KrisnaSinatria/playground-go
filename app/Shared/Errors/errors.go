package errors

import (
	"errors"
	"net/http"

	"go-first/app/Shared/Response"

	"github.com/gin-gonic/gin"
)

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type BusinessRuleError struct {
	Message string
}

func (e *BusinessRuleError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, err error) {
	var nfErr *NotFoundError
	var valErr *ValidationError
	var bizErr *BusinessRuleError
	var intErr *InternalServerError

	switch {
	case errors.As(err, &nfErr):
		response.ErrorJSON(c, http.StatusNotFound, nfErr.Error())
	case errors.As(err, &valErr):
		response.ErrorJSON(c, http.StatusBadRequest, valErr.Error())
	case errors.As(err, &bizErr):
		response.ErrorJSON(c, http.StatusUnprocessableEntity, bizErr.Error())
	case errors.As(err, &intErr):
		response.ErrorJSON(c, http.StatusInternalServerError, intErr.Error())
	default:
		response.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}
}
