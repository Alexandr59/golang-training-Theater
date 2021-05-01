package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type PlaceData struct {
	db *gorm.DB
}

func NewPlaceData(db *gorm.DB) *PlaceData {
	return &PlaceData{db: db}
}

type Place struct {
	Id       int    `gorm:"primaryKey"`
	SectorId int    `gorm:"sector_id"`
	Name     string `gorm:"name"`
}

func (p PlaceData) AddPlace(place Place) (int, error) {
	result := p.db.Create(&place)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Place to database, error: %w", result.Error)
	}
	return place.Id, nil
}

func (p PlaceData) DeletePlace(entry Place) error {
	result := p.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Place to database, error: %w", result.Error)
	}
	return nil
}

func (p PlaceData) UpdatePlace(place Place) error {
	result := p.db.Model(&place).Updates(place)
	if result.Error != nil {
		return fmt.Errorf("can't update Place to database, error: %w", result.Error)
	}
	return nil
}

func (p PlaceData) FindByIdPlace(entry Place) (Place, error) {
	result := p.db.First(&entry)
	if result.Error != nil {
		return Place{}, fmt.Errorf("can't find Place to database, error: %w", result.Error)
	}
	return entry, nil
}
