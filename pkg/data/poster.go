package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type PosterData struct {
	db *gorm.DB
}

func NewPosterData(db *gorm.DB) *PosterData {
	return &PosterData{db: db}
}

type Poster struct {
	Id         int    `gorm:"primaryKey"`
	AccountId  int    `gorm:"account_id"`
	ScheduleId int    `gorm:"schedule_id"`
	Comment    string `gorm:"comment"`
}

type SelectPoster struct {
	Id                  int
	PerformanceName     string
	GenreName           string
	PerformanceDuration string
	DateTime            string
	HallName            string
	HallCapacity        int
	LocationAddress     string
	LocationPhoneNumber string
	Comment             string
}

func (p PosterData) AddPoster(poster Poster) (int, error) {
	result := p.db.Create(&poster)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Poster to database, error: %w", result.Error)
	}
	return poster.Id, nil
}

func (p PosterData) DeletePoster(entry Poster) error {
	result := p.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Poster to database, error: %w", result.Error)
	}
	return nil
}

func (p PosterData) UpdatePoster(poster Poster) error {
	result := p.db.Model(&poster).Updates(poster)
	if result.Error != nil {
		return fmt.Errorf("can't update Poster to database, error: %w", result.Error)
	}
	return nil
}

func (p PosterData) FindByIdPoster(entry Poster) (Poster, error) {
	result := p.db.First(&entry)
	if result.Error != nil {
		return Poster{}, fmt.Errorf("can't find Poster to database, error: %w", result.Error)
	}
	return entry, nil
}

func (p PosterData) ReadAllPosters() ([]SelectPoster, error) {
	var posters []SelectPoster
	rows, err := p.db.Table("posters").Select("posters.id, performances.name, genres.name, " +
		"performances.duration, schedules.date, halls.name, halls.capacity, locations.address, locations.phone_number, posters.comment ").
		Joins("JOIN schedules on schedules.id = posters.schedule_id").
		Joins("JOIN performances on schedules.performance_id = performances.id").
		Joins("JOIN genres on performances.genre_id = genres.id").
		Joins("JOIN halls on schedules.hall_id = halls.id").
		Joins("JOIN locations on halls.location_id = locations.id").
		Rows()
	if err != nil {
		return nil, fmt.Errorf("can't get posters from database, error:%w", err)
	}
	for rows.Next() {
		var temp SelectPoster
		err = rows.Scan(&temp.Id, &temp.PerformanceName, &temp.GenreName, &temp.PerformanceDuration,
			&temp.DateTime, &temp.HallName, &temp.HallCapacity, &temp.LocationAddress, &temp.LocationPhoneNumber,
			&temp.Comment)
		if err != nil {
			return nil, fmt.Errorf("can't scan posters from database, error:%w", err)
		}
		posters = append(posters, temp)
	}
	return posters, nil
}
