package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type ScheduleData struct {
	db *gorm.DB
}

func NewScheduleData(db *gorm.DB) *ScheduleData {
	return &ScheduleData{db: db}
}

type Schedule struct {
	Id            int    `gorm:"primaryKey"`
	AccountId     int    `gorm:"account_id"`
	PerformanceId int    `gorm:"performance_id"`
	Date          string `gorm:"date"`
	HallId        int    `gorm:"hall_id"`
}

func (s ScheduleData) AddSchedule(schedule Schedule) (int, error) {
	result := s.db.Create(&schedule)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Schedule to database, error: %w", result.Error)
	}
	return schedule.Id, nil
}

func (s ScheduleData) DeleteSchedule(entry Schedule) error {
	result := s.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Schedule to database, error: %w", result.Error)
	}
	return nil
}

func (s ScheduleData) UpdateSchedule(schedule Schedule) error {
	result := s.db.Model(&schedule).Updates(schedule)
	if result.Error != nil {
		return fmt.Errorf("can't update Schedule to database, error: %w", result.Error)
	}
	return nil
}

func (s ScheduleData) FindByIdSchedule(entry Schedule) (Schedule, error) {
	result := s.db.First(&entry)
	if result.Error != nil {
		return Schedule{}, fmt.Errorf("can't find Schedule to database, error: %w", result.Error)
	}
	return entry, nil
}
