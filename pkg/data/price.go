package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type PriceData struct {
	db *gorm.DB
}

func NewPriceData(db *gorm.DB) *PriceData {
	return &PriceData{db: db}
}

type Price struct {
	Id            int `gorm:"primaryKey"`
	AccountId     int `gorm:"account_id"`
	SectorId      int `gorm:"sector_id"`
	PerformanceId int `gorm:"performance_id"`
	Price         int `gorm:"price"`
}

func (p PriceData) AddPrice(price Price) (int, error) {
	result := p.db.Create(&price)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Price to database, error: %w", result.Error)
	}
	return price.Id, nil
}

func (p PriceData) DeletePrice(entry Price) error {
	result := p.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Price to database, error: %w", result.Error)
	}
	return nil
}

func (p PriceData) UpdatePrice(price Price) error {
	result := p.db.Model(&price).Updates(price)
	if result.Error != nil {
		return fmt.Errorf("can't update Price to database, error: %w", result.Error)
	}
	return nil
}

func (p PriceData) FindByIdPrice(entry Price) (Price, error) {
	result := p.db.First(&entry)
	if result.Error != nil {
		return Price{}, fmt.Errorf("can't find Price to database, error: %w", result.Error)
	}
	return entry, nil
}
