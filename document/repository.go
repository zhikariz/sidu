package document

import (
	. "sidu/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Save(document Document) (Document, error)
	FindAll() ([]Document, error)
	FindById(id int) (Document, error)
	Update(document Document) (Document, error)
	Delete(document Document) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(document Document) (Document, error) {
	err := r.db.Create(&document).Error

	if err != nil {
		return document, err
	}

	return document, nil
}

func (r *repository) FindById(id int) (Document, error) {
	var document Document
	err := r.db.Where("id = ?", uint(id)).Find(&document).Error

	if err != nil {
		return document, err
	}
	return document, nil
}

func (r *repository) Update(document Document) (Document, error) {
	err := r.db.Save(&document).Error
	if err != nil {
		return document, err
	}
	return document, nil
}

func (r *repository) Delete(document Document) (bool, error) {
	err := r.db.Delete(&document).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindAll() ([]Document, error) {
	var documents []Document
	err := r.db.Find(&documents).Error
	if err != nil {
		return documents, err
	}
	return documents, nil
}
