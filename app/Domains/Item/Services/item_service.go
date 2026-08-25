package services

import (
	"go-first/app/Domains/Item/DTO"
	"go-first/app/Domains/Item/Models"
	appErrors "go-first/app/Shared/Errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{db: db}
}

func (s *ItemService) Create(req dto.CreateItemRequest) (*dto.ItemResponse, error) {
	item := models.Item{
		Name:           req.Name,
		SKU:            req.SKU,
		QuantityOnHand: 0,
		SellingPrice:   req.SellingPrice,
	}

	if err := s.db.Create(&item).Error; err != nil {
		return nil, &appErrors.InternalServerError{Message: "Gagal membuat item: " + err.Error()}
	}

	return s.toResponse(&item), nil
}

func (s *ItemService) GetAll(page, limit int, search string) ([]dto.ItemResponse, int64, error) {
	var items []models.Item
	var total int64

	db := s.db.Model(&models.Item{})

	if search != "" {
		db = db.Where("name ILIKE ? OR sku ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal menghitung total item"}
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		db = db.Offset(offset).Limit(limit)
	}

	if err := db.Find(&items).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal mengambil daftar item"}
	}

	var responses []dto.ItemResponse
	for _, item := range items {
		responses = append(responses, *s.toResponse(&item))
	}

	return responses, total, nil
}

func (s *ItemService) GetByID(id uint) (*dto.ItemResponse, error) {
	var item models.Item
	if err := s.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.NotFoundError{Message: "Item tidak ditemukan"}
		}
		return nil, &appErrors.InternalServerError{Message: err.Error()}
	}

	return s.toResponse(&item), nil
}

func (s *ItemService) Update(id uint, req dto.UpdateItemRequest) (*dto.ItemResponse, error) {
	var item models.Item
	if err := s.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.NotFoundError{Message: "Item tidak ditemukan"}
		}
		return nil, &appErrors.InternalServerError{Message: err.Error()}
	}

	item.Name = req.Name
	item.SKU = req.SKU
	item.SellingPrice = req.SellingPrice

	if err := s.db.Save(&item).Error; err != nil {
		return nil, &appErrors.InternalServerError{Message: "Gagal mengupdate item"}
	}

	return s.toResponse(&item), nil
}

func (s *ItemService) Delete(id uint) error {
	var item models.Item
	if err := s.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &appErrors.NotFoundError{Message: "Item tidak ditemukan"}
		}
		return &appErrors.InternalServerError{Message: err.Error()}
	}

	if err := s.db.Delete(&item).Error; err != nil {
		return &appErrors.InternalServerError{Message: "Gagal menghapus item"}
	}

	return nil
}

func (s *ItemService) IncrementStock(tx *gorm.DB, itemID uint, qty int) error {
	result := tx.Model(&models.Item{}).
		Where("id = ?", itemID).
		UpdateColumn("quantity_on_hand", gorm.Expr("quantity_on_hand + ?", qty))

	if result.Error != nil {
		return &appErrors.InternalServerError{Message: "Gagal menambah stok item: " + result.Error.Error()}
	}

	if result.RowsAffected == 0 {
		return &appErrors.NotFoundError{Message: "Item tidak ditemukan"}
	}

	return nil
}

func (s *ItemService) DecrementStockWithLock(tx *gorm.DB, itemID uint, qty int) (float64, error) {
	var item models.Item

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", itemID).
		First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, &appErrors.NotFoundError{Message: "Item tidak ditemukan"}
		}
		return 0, &appErrors.InternalServerError{Message: "Gagal mengunci data item: " + err.Error()}
	}

	if item.QuantityOnHand < qty {
		return 0, &appErrors.BusinessRuleError{Message: "Stok tidak mencukupi untuk item: " + item.Name}
	}

	if err := tx.Model(&item).
		UpdateColumn("quantity_on_hand", gorm.Expr("quantity_on_hand - ?", qty)).Error; err != nil {
		return 0, &appErrors.InternalServerError{Message: "Gagal mengurangkan stok item: " + err.Error()}
	}

	return item.SellingPrice, nil
}

func (s *ItemService) toResponse(item *models.Item) *dto.ItemResponse {
	return &dto.ItemResponse{
		ID:             item.ID,
		Name:           item.Name,
		SKU:            item.SKU,
		QuantityOnHand: item.QuantityOnHand,
		SellingPrice:   item.SellingPrice,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
	}
}
