package controllers

import (
	"net/http"
	"strconv"

	"go-first/app/Domains/Technician/DTO"
	"go-first/app/Domains/Technician/Services"
	appErrors "go-first/app/Shared/Errors"
	"go-first/app/Shared/Response"

	"github.com/gin-gonic/gin"
)

type TechnicianController struct {
	service *services.TechnicianService
}

func NewTechnicianController(service *services.TechnicianService) *TechnicianController {
	return &TechnicianController{service: service}
}

func (ctrl *TechnicianController) Create(c *gin.Context) {
	var req dto.CreateTechnicianRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.Create(req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusCreated, "Teknisi berhasil dibuat", res)
}

func (ctrl *TechnicianController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	res, total, err := ctrl.service.GetAll(page, limit)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.PaginatedJSON(c, http.StatusOK, "Berhasil mengambil daftar teknisi", res, page, limit, total)
}

func (ctrl *TechnicianController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID teknisi tidak valid"})
		return
	}

	res, err := ctrl.service.GetByID(uint(id))
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Berhasil mengambil detail teknisi", res)
}

func (ctrl *TechnicianController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID teknisi tidak valid"})
		return
	}

	var req dto.UpdateTechnicianRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.Update(uint(id), req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Teknisi berhasil diperbarui", res)
}

func (ctrl *TechnicianController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID teknisi tidak valid"})
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Teknisi berhasil dihapus", nil)
}
