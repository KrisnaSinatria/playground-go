package response

import (
	"math"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

type PaginatedSuccessResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func JSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func PaginatedJSON(c *gin.Context, statusCode int, message string, data interface{}, page, limit int, totalItems int64) {
	totalPages := 0
	if limit > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(limit)))
	}
	c.JSON(statusCode, PaginatedSuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta: PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	})
}

func ErrorJSON(c *gin.Context, statusCode int, errMessage string) {
	c.JSON(statusCode, ErrorResponse{
		Success: false,
		Error:   errMessage,
	})
}
