package controllers

import (
	"net/http"
	"strconv"

	"go-first/app/Domains/StockMovement/DTO"
	"go-first/app/Domains/StockMovement/Services"
	appErrors "go-first/app/Shared/Errors"
	"go-first/app/Shared/Response"

	"github.com/gin-gonic/gin"
)

type StockMovementController struct {
	service *services.StockMovementService
}

func NewStockMovementController(service *services.StockMovementService) *StockMovementController {
	return &StockMovementController{service: service}
}

func (ctrl *StockMovementController) StockIn(c *gin.Context) {
	var req dto.StockInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.StockIn(req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusCreated, "Stok masuk (Restock) berhasil dicatat", res)
}

func (ctrl *StockMovementController) StockOut(c *gin.Context) {
	var req dto.StockOutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.StockOut(req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusCreated, "Stok keluar (Pemakaian repair) berhasil dicatat", res)
}

func (ctrl *StockMovementController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	movementType := c.Query("type")

	var itemIDPtr *uint
	if itemIDStr := c.Query("item_id"); itemIDStr != "" {
		if id, err := strconv.ParseUint(itemIDStr, 10, 64); err == nil {
			uID := uint(id)
			itemIDPtr = &uID
		}
	}

	var repairIDPtr *uint
	if repairIDStr := c.Query("repair_id"); repairIDStr != "" {
		if id, err := strconv.ParseUint(repairIDStr, 10, 64); err == nil {
			uID := uint(id)
			repairIDPtr = &uID
		}
	}

	res, total, err := ctrl.service.GetAll(page, limit, itemIDPtr, repairIDPtr, movementType)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.PaginatedJSON(c, http.StatusOK, "Berhasil mengambil riwayat pergerakan stok", res, page, limit, total)
}
