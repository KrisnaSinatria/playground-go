package router

import (
	itemControllers "go-first/app/Domains/Item/Controllers"
	itemServices "go-first/app/Domains/Item/Services"

	repairControllers "go-first/app/Domains/Repair/Controllers"
	repairServices "go-first/app/Domains/Repair/Services"

	stockControllers "go-first/app/Domains/StockMovement/Controllers"
	stockServices "go-first/app/Domains/StockMovement/Services"

	techControllers "go-first/app/Domains/Technician/Controllers"
	techServices "go-first/app/Domains/Technician/Services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	techService := techServices.NewTechnicianService(db)
	itemService := itemServices.NewItemService(db)
	stockService := stockServices.NewStockMovementService(db, itemService)
	repairService := repairServices.NewRepairService(db, techService, stockService)

	techCtrl := techControllers.NewTechnicianController(techService)
	itemCtrl := itemControllers.NewItemController(itemService)
	stockCtrl := stockControllers.NewStockMovementController(stockService)
	repairCtrl := repairControllers.NewRepairController(repairService)

	techGroup := r.Group("/technicians")
	{
		techGroup.POST("", techCtrl.Create)
		techGroup.GET("", techCtrl.GetAll)
		techGroup.GET("/:id", techCtrl.GetByID)
		techGroup.PUT("/:id", techCtrl.Update)
		techGroup.DELETE("/:id", techCtrl.Delete)
	}

	itemGroup := r.Group("/items")
	{
		itemGroup.POST("", itemCtrl.Create)
		itemGroup.GET("", itemCtrl.GetAll)
		itemGroup.GET("/:id", itemCtrl.GetByID)
		itemGroup.PUT("/:id", itemCtrl.Update)
		itemGroup.DELETE("/:id", itemCtrl.Delete)
	}

	stockGroup := r.Group("/stock-movements")
	{
		stockGroup.POST("/stock-in", stockCtrl.StockIn)
		stockGroup.POST("/stock-out", stockCtrl.StockOut)
		stockGroup.GET("", stockCtrl.GetAll)
	}

	repairGroup := r.Group("/repairs")
	{
		repairGroup.POST("", repairCtrl.Create)
		repairGroup.GET("", repairCtrl.GetAll)
		repairGroup.GET("/:id", repairCtrl.GetByID)
		repairGroup.PUT("/:id/status", repairCtrl.UpdateStatus)
	}

	return r
}
