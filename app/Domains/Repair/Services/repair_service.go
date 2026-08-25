package services

import (
	"go-first/app/Domains/Repair/DTO"
	"go-first/app/Domains/Repair/Models"
	stockDTO "go-first/app/Domains/StockMovement/DTO"
	stockMovementServices "go-first/app/Domains/StockMovement/Services"
	technicianServices "go-first/app/Domains/Technician/Services"
	appErrors "go-first/app/Shared/Errors"

	"gorm.io/gorm"
)

type RepairService struct {
	db                   *gorm.DB
	technicianService    *technicianServices.TechnicianService
	stockMovementService *stockMovementServices.StockMovementService
}

func NewRepairService(
	db *gorm.DB,
	technicianService *technicianServices.TechnicianService,
	stockMovementService *stockMovementServices.StockMovementService,
) *RepairService {
	return &RepairService{
		db:                   db,
		technicianService:    technicianService,
		stockMovementService: stockMovementService,
	}
}

func (s *RepairService) Create(req dto.CreateRepairRequest) (*dto.RepairResponse, error) {
	_, err := s.technicianService.GetByID(req.TechnicianID)
	if err != nil {
		return nil, err
	}

	repair := models.Repair{
		TechnicianID: req.TechnicianID,
		NameCustomer: req.NameCustomer,
		Description:  req.Description,
		ServiceFee:   req.ServiceFee,
		Status:       models.StatusPending,
	}

	if err := s.db.Create(&repair).Error; err != nil {
		return nil, &appErrors.InternalServerError{Message: "Gagal membuat tiket repair: " + err.Error()}
	}

	return s.toResponse(&repair), nil
}

func (s *RepairService) GetAll(page, limit int, status string, technicianID *uint) ([]dto.RepairResponse, int64, error) {
	var repairs []models.Repair
	var total int64

	db := s.db.Model(&models.Repair{})

	if status != "" {
		db = db.Where("status = ?", status)
	}
	if technicianID != nil {
		db = db.Where("technician_id = ?", *technicianID)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal menghitung total tiket repair"}
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		db = db.Order("created_at desc").Offset(offset).Limit(limit)
	} else {
		db = db.Order("created_at desc")
	}

	if err := db.Find(&repairs).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal mengambil daftar tiket repair"}
	}

	var responses []dto.RepairResponse
	for _, r := range repairs {
		responses = append(responses, *s.toResponse(&r))
	}

	return responses, total, nil
}

func (s *RepairService) GetByID(id uint) (*dto.RepairDetailResponse, error) {
	var repair models.Repair
	if err := s.db.First(&repair, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.NotFoundError{Message: "Tiket repair tidak ditemukan"}
		}
		return nil, &appErrors.InternalServerError{Message: err.Error()}
	}

	techName := ""
	if tech, err := s.technicianService.GetByID(repair.TechnicianID); err == nil {
		techName = tech.Name
	}

	usedItems, err := s.stockMovementService.GetByRepairID(repair.ID)
	if err != nil {
		usedItems = []stockDTO.StockMovementResponse{}
	}

	var itemsTotalCost float64
	for _, item := range usedItems {
		itemsTotalCost += (item.UnitPrice * float64(item.Quantity))
	}
	totalCost := repair.ServiceFee + itemsTotalCost

	return &dto.RepairDetailResponse{
		ID:             repair.ID,
		TechnicianID:   repair.TechnicianID,
		TechnicianName: techName,
		NameCustomer:   repair.NameCustomer,
		Description:    repair.Description,
		ServiceFee:     repair.ServiceFee,
		Status:         repair.Status,
		UsedItems:      usedItems,
		TotalCost:      totalCost,
		CreatedAt:      repair.CreatedAt,
		UpdatedAt:      repair.UpdatedAt,
	}, nil
}

func (s *RepairService) UpdateStatus(id uint, req dto.UpdateRepairStatusRequest) (*dto.RepairResponse, error) {
	var repair models.Repair
	if err := s.db.First(&repair, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.NotFoundError{Message: "Tiket repair tidak ditemukan"}
		}
		return nil, &appErrors.InternalServerError{Message: err.Error()}
	}

	repair.Status = req.Status

	if err := s.db.Save(&repair).Error; err != nil {
		return nil, &appErrors.InternalServerError{Message: "Gagal memperbarui status repair"}
	}

	return s.toResponse(&repair), nil
}

func (s *RepairService) toResponse(r *models.Repair) *dto.RepairResponse {
	return &dto.RepairResponse{
		ID:           r.ID,
		TechnicianID: r.TechnicianID,
		NameCustomer: r.NameCustomer,
		Description:  r.Description,
		ServiceFee:   r.ServiceFee,
		Status:       r.Status,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}
