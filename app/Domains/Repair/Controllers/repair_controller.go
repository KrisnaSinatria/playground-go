package controllers

import (
	"net/http"
	"strconv"

	"go-first/app/Domains/Repair/DTO"
	"go-first/app/Domains/Repair/Services"
	appErrors "go-first/app/Shared/Errors"
	"go-first/app/Shared/Response"

	"github.com/gin-gonic/gin"
)

type RepairController struct {
	service *services.RepairService
}

func NewRepairController(service *services.RepairService) *RepairController {
	return &RepairController{service: service}
}

func (ctrl *RepairController) Create(c *gin.Context) {
	var req dto.CreateRepairRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.Create(req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusCreated, "Tiket repair berhasil dibuat", res)
}

func (ctrl *RepairController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.Query("status")

	var techIDPtr *uint
	if techIDStr := c.Query("technician_id"); techIDStr != "" {
		if id, err := strconv.ParseUint(techIDStr, 10, 64); err == nil {
			uID := uint(id)
			techIDPtr = &uID
		}
	}

	res, total, err := ctrl.service.GetAll(page, limit, status, techIDPtr)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.PaginatedJSON(c, http.StatusOK, "Berhasil mengambil daftar tiket repair", res, page, limit, total)
}

func (ctrl *RepairController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID tiket repair tidak valid"})
		return
	}

	res, err := ctrl.service.GetByID(uint(id))
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Berhasil mengambil detail tiket repair", res)
}

func (ctrl *RepairController) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID tiket repair tidak valid"})
		return
	}

	var req dto.UpdateRepairStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.UpdateStatus(uint(id), req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Status tiket repair berhasil diperbarui", res)
}
