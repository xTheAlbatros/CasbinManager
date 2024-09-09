package repositories

import (
	"CasbinManager/models"
	"gorm.io/gorm"
)

type CasbinRepository struct {
	db *gorm.DB
}

func NewCasbinRepository(db *gorm.DB) *CasbinRepository {
	return &CasbinRepository{db: db}
}

func (r *CasbinRepository) GetAllRules() ([]models.CasbinRule, error) {
	var rules []models.CasbinRule
	err := r.db.Table("public.casbin_rule").Find(&rules).Error
	return rules, err
}

func (r *CasbinRepository) AddRule(rule models.CasbinRule) error {
	return r.db.Table("public.casbin_rule").Create(&rule).Error
}

func (r *CasbinRepository) DeleteRule(id int) error {
	return r.db.Table("public.casbin_rule").Delete(&models.CasbinRule{}, id).Error
}

func (r *CasbinRepository) UpdateRule(rule models.CasbinRule) error {
	return r.db.Table("public.casbin_rule").Where("id = ?", rule.ID).Updates(&rule).Error
}
