package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type HallData struct {
	db *gorm.DB
}

func NewHallData(db *gorm.DB) *HallData {
	return &HallData{db: db}
}

type Hall struct {
	Id         int    `gorm:"primaryKey"`
	AccountId  int    `gorm:"account_id"`
	Name       string `gorm:"name"`
	Capacity   int    `gorm:"capacity"`
	LocationId int    `gorm:"location_id"`
}

func (h HallData) AddHall(hall Hall) (int, error) {
	result := h.db.Create(&hall)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert hall to database, error: %w", result.Error)
	}
	return hall.Id, nil
}

func (h HallData) DeleteHall(entry Hall) error {
	result := h.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Hall to database, error: %w", result.Error)
	}
	return nil
}

func (h HallData) UpdateHall(hall Hall) error {
	result := h.db.Model(&hall).Updates(hall)
	if result.Error != nil {
		return fmt.Errorf("can't update hall to database, error: %w", result.Error)
	}
	return nil
}

func (h HallData) FindByIdHall(entry Hall) (Hall, error) {
	result := h.db.First(&entry)
	if result.Error != nil {
		return Hall{}, fmt.Errorf("can't find Hall to database, error: %w", result.Error)
	}
	return entry, nil
}
