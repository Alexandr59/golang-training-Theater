package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type PerformanceData struct {
	db *gorm.DB
}

func NewPerformanceData(db *gorm.DB) *PerformanceData {
	return &PerformanceData{db: db}
}

type Performance struct {
	Id        int    `gorm:"primaryKey"`
	AccountId int    `gorm:"account_id"`
	Name      string `gorm:"name"`
	GenreId   int    `gorm:"genre_id"`
	Duration  string `gorm:"duration"`
}

func (p PerformanceData) AddPerformance(performance Performance) (int, error) {
	result := p.db.Create(&performance)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Performance to database, error: %w", result.Error)
	}
	return performance.Id, nil
}

func (p PerformanceData) DeletePerformance(entry Performance) error {
	result := p.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Performance to database, error: %w", result.Error)
	}
	return nil
}

func (p PerformanceData) UpdatePerformance(performance Performance) error {
	result := p.db.Model(&performance).Updates(performance)
	if result.Error != nil {
		return fmt.Errorf("can't update Performance to database, error: %w", result.Error)
	}
	return nil
}

func (p PerformanceData) FindByIdPerformance(entry Performance) (Performance, error) {
	result := p.db.First(&entry)
	if result.Error != nil {
		return Performance{}, fmt.Errorf("can't find Performance to database, error: %w", result.Error)
	}
	return entry, nil
}
