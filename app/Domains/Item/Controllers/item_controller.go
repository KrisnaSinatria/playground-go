package controllers

import (
	"net/http"
	"strconv"

	"go-first/app/Domains/Item/DTO"
	"go-first/app/Domains/Item/Services"
	appErrors "go-first/app/Shared/Errors"
	"go-first/app/Shared/Response"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	service *services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (ctrl *ItemController) Create(c *gin.Context) {
	var req dto.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.Create(req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusCreated, "Item berhasil dibuat", res)
}

func (ctrl *ItemController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	res, total, err := ctrl.service.GetAll(page, limit, search)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.PaginatedJSON(c, http.StatusOK, "Berhasil mengambil daftar item", res, page, limit, total)
}

func (ctrl *ItemController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID item tidak valid"})
		return
	}

	res, err := ctrl.service.GetByID(uint(id))
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Berhasil mengambil detail item", res)
}

func (ctrl *ItemController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID item tidak valid"})
		return
	}

	var req dto.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: err.Error()})
		return
	}

	res, err := ctrl.service.Update(uint(id), req)
	if err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Item berhasil diperbarui", res)
}

func (ctrl *ItemController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		appErrors.HandleError(c, &appErrors.ValidationError{Message: "ID item tidak valid"})
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		appErrors.HandleError(c, err)
		return
	}

	response.JSON(c, http.StatusOK, "Item berhasil dihapus", nil)
}
