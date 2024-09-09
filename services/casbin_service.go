package services

import (
	"CasbinManager/models"
	"CasbinManager/repositories"
	"gorm.io/gorm"
)

type CasbinService struct {
	repo *repositories.CasbinRepository
}

func NewCasbinService(db *gorm.DB) *CasbinService {
	return &CasbinService{
		repo: repositories.NewCasbinRepository(db),
	}
}

func (s *CasbinService) GetAllRules() ([]models.CasbinRule, error) {
	return s.repo.GetAllRules()
}

func (s *CasbinService) AddRule(rule models.CasbinRule) error {
	return s.repo.AddRule(rule)
}

func (s *CasbinService) DeleteRule(id int) error {
	return s.repo.DeleteRule(id)
}

func (s *CasbinService) UpdateRule(rule models.CasbinRule) error {
	return s.repo.UpdateRule(rule)
}
