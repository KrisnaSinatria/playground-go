package services

import (
	"go-first/app/Domains/Technician/DTO"
	"go-first/app/Domains/Technician/Models"
	appErrors "go-first/app/Shared/Errors"

	"gorm.io/gorm"
)

type TechnicianService struct {
	db *gorm.DB
}

func NewTechnicianService(db *gorm.DB) *TechnicianService {
	return &TechnicianService{db: db}
}

func (s *TechnicianService) Create(req dto.CreateTechnicianRequest) (*dto.TechnicianResponse, error) {
	tech := models.Technician{
		Name:  req.Name,
		Phone: req.Phone,
	}

	if err := s.db.Create(&tech).Error; err != nil {
		return nil, &appErrors.InternalServerError{Message: "Gagal membuat teknisi: " + err.Error()}
	}

	return s.toResponse(&tech), nil
}

func (s *TechnicianService) GetAll(page, limit int) ([]dto.TechnicianResponse, int64, error) {
	var technicians []models.Technician
	var total int64

	db := s.db.Model(&models.Technician{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal menghitung total teknisi"}
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		db = db.Offset(offset).Limit(limit)
	}

	if err := db.Find(&technicians).Error; err != nil {
		return nil, 0, &appErrors.InternalServerError{Message: "Gagal mengambil daftar teknisi"}
	}

	var responses []dto.TechnicianResponse
	for _, tech := range technicians {
		responses = append(responses, *s.toResponse(&tech))
	}

	return responses, total, nil
}

func (s *TechnicianService) GetByID(id uint) (*dto.TechnicianResponse, error) {
	var tech models.Technician
	if err := s.db.First(&tech, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.NotFoundError{Message: "Teknisi tidak ditemukan"}
		}
		return nil, &appErrors.InternalServerError{Message: err.Error()}
	}

	return s.toResponse(&tech), nil
}

func (s *TechnicianService) Update(id uint, req dto.UpdateTechnicianRequest) (*dto.TechnicianResponse, error) {
	var tech models.Technician
	if err := s.db.First(&tech, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &appErrors.NotFoundError{Message: "Teknisi tidak ditemukan"}
		}
		return nil, &appErrors.InternalServerError{Message: err.Error()}
	}

	tech.Name = req.Name
	tech.Phone = req.Phone

	if err := s.db.Save(&tech).Error; err != nil {
		return nil, &appErrors.InternalServerError{Message: "Gagal mengupdate teknisi"}
	}

	return s.toResponse(&tech), nil
}

func (s *TechnicianService) Delete(id uint) error {
	var tech models.Technician
	if err := s.db.First(&tech, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &appErrors.NotFoundError{Message: "Teknisi tidak ditemukan"}
		}
		return &appErrors.InternalServerError{Message: err.Error()}
	}

	if err := s.db.Delete(&tech).Error; err != nil {
		return &appErrors.InternalServerError{Message: "Gagal menghapus teknisi"}
	}

	return nil
}

func (s *TechnicianService) toResponse(tech *models.Technician) *dto.TechnicianResponse {
	return &dto.TechnicianResponse{
		ID:        tech.ID,
		Name:      tech.Name,
		Phone:     tech.Phone,
		CreatedAt: tech.CreatedAt,
		UpdatedAt: tech.UpdatedAt,
	}
}
