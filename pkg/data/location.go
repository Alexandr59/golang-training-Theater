package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type LocationData struct {
	db *gorm.DB
}

func NewLocationData(db *gorm.DB) *LocationData {
	return &LocationData{db: db}
}

type Location struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	Address     string `gorm:"address"`
	PhoneNumber string `gorm:"phone_number"`
}

func (l LocationData) AddLocation(location Location) (int, error) {
	result := l.db.Create(&location)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert location to database, error: %w", result.Error)
	}
	return location.Id, nil
}

func (l LocationData) DeleteLocation(entry Location) error {
	result := l.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Location to database, error: %w", result.Error)
	}
	return nil
}

func (l LocationData) UpdateLocation(location Location) error {
	result := l.db.Model(&location).Updates(location)
	if result.Error != nil {
		return fmt.Errorf("can't update location to database, error: %w", result.Error)
	}
	return nil
}

func (l LocationData) FindByIdLocation(entry Location) (Location, error) {
	result := l.db.First(&entry)
	if result.Error != nil {
		return Location{}, fmt.Errorf("can't find Location to database, error: %w", result.Error)
	}
	return entry, nil
}
