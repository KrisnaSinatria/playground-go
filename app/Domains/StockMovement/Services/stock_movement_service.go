package services

import (
	itemServices "go-first/app/Domains/Item/Services"
	"go-first/app/Domains/StockMovement/DTO"
	"go-first/app/Domains/StockMovement/Models"
	appErrors "go-first/app/Shared/Errors"

	"gorm.io/gorm"
)

type StockMovementService struct {
	db          *gorm.DB
	itemService *itemServices.ItemService
}

func NewStockMovementService(db *gorm.DB, itemService *itemServices.ItemService) *StockMovementService {
	return &StockMovementService{
		db:          db,
		itemService: itemService,
	}
}

func (s *StockMovementService) StockIn(req dto.StockInRequest) (*dto.StockMovementResponse, error) {
	var movement models.StockMovement

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.itemService.IncrementStock(tx, req.ItemID, req.Quantity); err != nil {
			return err
		}

		movement = models.StockMovement{
			ItemID:    req.ItemID,
			Type:      models.MovementIn,
			Quantity:  req.Quantity,
			UnitPrice: req.UnitPrice,
			Note:      req.Note,
		}

		if err := tx.Create(&movement).Error; err != nil {
			return &appErrors.InternalServerError{Message: "Gagal mencatat pergerakan stok: " + err.Error()}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.toResponse(&movement), nil
}

func (s *StockMovementService) StockOut(req dto.StockOutRequest) (*dto.StockMovementResponse, error) {
	var movement models.StockMovement

	err := s.db.Transaction(func(tx *gorm.DB) error {
		sellingPrice, err := s.itemService.DecrementStockWithLock(tx, req.ItemID, req.Quantity)
		if err != nil {
			return err
		}

		repairID := req.RepairID
		movement = models.StockMovement{
			ItemID:    req.ItemID,
			RepairID:  &repairID,
			Type:      models.MovementOut,
			Quantity:  req.Quantity,
			UnitPrice: sellingPrice,
			Note:      req.Note,
		}

		if err := tx.Create(&movement).Error; err != nil {
			return &appErrors.InternalServerError{Message: "Gagal mencatat pergerakan stok keluar: " + err.Error()}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.toResponse(&movement), nil
}

func (s *StockMovementService) GetAll(page, limit int, itemID, repairID *uint, movementType string) ([]dto.StockMovementResponse, int64, error) {
	var movements []models.StockMovement
	var total int64

	db := s.db.Model(&models.StockMovement{})

	if itemID != nil {
		db = db.Where("item_id = ?", *itemID)
	}
	if repairID != nil {
		db = db.Where("repair_id = ?", *repairID)
	}
	if movementType != "" {
		db = db.Where("type = ?", movementType)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal menghitung total pergerakan stok"}
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		db = db.Order("created_at desc").Offset(offset).Limit(limit)
	} else {
		db = db.Order("created_at desc")
	}

	if err := db.Find(&movements).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal mengambil daftar pergerakan stok"}
	}

	var responses []dto.StockMovementResponse
	for _, m := range movements {
		responses = append(responses, *s.toResponse(&m))
	}

	return responses, total, nil
}

func (s *StockMovementService) GetByRepairID(repairID uint) ([]dto.StockMovementResponse, error) {
	var movements []models.StockMovement
	if err := s.db.Where("repair_id = ?", repairID).Find(&movements).Error; err != nil {
		return nil, &appErrors.InternalServerError{Message: "Gagal mengambil pergerakan stok repair"}
	}

	var responses []dto.StockMovementResponse
	for _, m := range movements {
		responses = append(responses, *s.toResponse(&m))
	}

	return responses, nil
}

func (s *StockMovementService) toResponse(m *models.StockMovement) *dto.StockMovementResponse {
	return &dto.StockMovementResponse{
		ID:        m.ID,
		ItemID:    m.ItemID,
		RepairID:  m.RepairID,
		Type:      m.Type,
		Quantity:  m.Quantity,
		UnitPrice: m.UnitPrice,
		Note:      m.Note,
		CreatedAt: m.CreatedAt,
	}
}
