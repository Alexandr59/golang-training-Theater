package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type GenreData struct {
	db *gorm.DB
}

func NewGenreData(db *gorm.DB) *GenreData {
	return &GenreData{db: db}
}

type Genre struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

func (g GenreData) AddGenre(genre Genre) (int, error) {
	result := g.db.Create(&genre)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert genre to database, error: %w", result.Error)
	}
	return genre.Id, nil
}

func (g GenreData) DeleteGenre(entry Genre) error {
	result := g.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Genre to database, error: %w", result.Error)
	}
	return nil
}

func (g GenreData) UpdateGenre(genre Genre) error {
	result := g.db.Model(&genre).Updates(genre)
	if result.Error != nil {
		return fmt.Errorf("can't update genre to database, error: %w", result.Error)
	}
	return nil
}

func (g GenreData) FindByIdGenre(entry Genre) (Genre, error) {
	result := g.db.First(&entry)
	if result.Error != nil {
		return Genre{}, fmt.Errorf("can't find Genre to database, error: %w", result.Error)
	}
	return entry, nil
}
