package model

import "gorm.io/gorm"

type GormBookModel struct {
	db *gorm.DB
}


//MUST ASK
func (m *GormBookModel) Get() []Book {
	var books []Book
	m.db.Find(&books)
	return books
}

func (m *GormBookModel) Insert(book Book) error {
	tx := m.db.Save(&book)
	return tx.Error
}

func NewGormBookModel(db *gorm.DB) *GormBookModel {
	return &GormBookModel{db: db}
}
