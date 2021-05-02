package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type SectorData struct {
	db *gorm.DB
}

func NewSectorData(db *gorm.DB) *SectorData {
	return &SectorData{db: db}
}

type Sector struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

func (s SectorData) AddSector(sector Sector) (int, error) {
	result := s.db.Create(&sector)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Sector to database, error: %w", result.Error)
	}
	return sector.Id, nil
}

func (s SectorData) DeleteSector(entry Sector) error {
	result := s.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Sector to database, error: %w", result.Error)
	}
	return nil
}

func (s SectorData) UpdateSector(sector Sector) error {
	result := s.db.Model(&sector).Updates(sector)
	if result.Error != nil {
		return fmt.Errorf("can't update Sector to database, error: %w", result.Error)
	}
	return nil
}

func (s SectorData) FindByIdSector(entry Sector) (Sector, error) {
	result := s.db.First(&entry)
	if result.Error != nil {
		return Sector{}, fmt.Errorf("can't find Sector to database, error: %w", result.Error)
	}
	return entry, nil
}
