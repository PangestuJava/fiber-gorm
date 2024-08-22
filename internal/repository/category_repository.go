package repository

import (
	"fiber-gorm/internal/domain/category"
	infrastructure "fiber-gorm/internal/infrastructure/database"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() category.Repository {
	return &categoryRepository{
		db: infrastructure.DB,
	}
}

func (r *categoryRepository) FindAll() ([]category.Category, error) {
	var categories []category.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindByID(id uint) (*category.Category, error) {
	var category category.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *category.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *category.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&category.Category{}, id).Error
}
