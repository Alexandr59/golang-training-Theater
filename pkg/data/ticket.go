package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type TicketData struct {
	db *gorm.DB
}

func NewTicketData(db *gorm.DB) *TicketData {
	return &TicketData{db: db}
}

type Ticket struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	ScheduleId  int    `gorm:"schedule_id"`
	PlaceId     int    `gorm:"place_id"`
	DateOfIssue string `gorm:"date_of_issue"`
	Paid        bool   `gorm:"paid"`
	Reservation bool   `gorm:"reservation"`
	Destroyed   bool   `gorm:"destroyed"`
}

type SelectTicket struct {
	Id                  int
	PerformanceName     string
	GenreName           string
	PerformanceDuration string
	DateTime            string
	HallName            string
	HallCapacity        int
	LocationAddress     string
	LocationPhoneNumber string
	SectorName          string
	Place               int
	Price               int
	DateOfIssue         string
	Paid                bool
	Reservation         bool
	Destroyed           bool
}

func (t TicketData) AddTicket(ticket Ticket) (int, error) {
	result := t.db.Create(&ticket)
	if result.Error != nil {
		return -1, fmt.Errorf("can't insert Ticket to database, error: %w", result.Error)
	}
	return ticket.Id, nil
}

func (t TicketData) DeleteTicket(entry Ticket) error {
	result := t.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Ticket to database, error: %w", result.Error)
	}
	return nil
}

func (t TicketData) UpdateTicket(ticket Ticket) error {
	result := t.db.Model(&ticket).Updates(ticket)
	if result.Error != nil {
		return fmt.Errorf("can't update Ticket to database, error: %w", result.Error)
	}
	return nil
}

func (t TicketData) FindByIdTicket(entry Ticket) (Ticket, error) {
	result := t.db.First(&entry)
	if result.Error != nil {
		return Ticket{}, fmt.Errorf("can't find Ticket to database, error: %w", result.Error)
	}
	return entry, nil
}

func (t TicketData) ReadAllTickets() ([]SelectTicket, error) {
	var tickets []SelectTicket
	rows, err := t.db.Table("tickets").Select("tickets.id, performances.name, genres.name, " +
		"performances.duration, schedules.date, halls.name, halls.capacity, locations.address, " +
		"locations.phone_number, sectors.name, places.name, prices.price, tickets.date_of_issue, " +
		"tickets.paid, tickets.reservation, tickets.destroyed").
		Joins("JOIN schedules on schedules.id = tickets.schedule_id").
		Joins("JOIN performances on schedules.performance_id = performances.id").
		Joins("JOIN genres on performances.genre_id = genres.id").
		Joins("JOIN halls on schedules.hall_id = halls.id").
		Joins("JOIN locations on halls.location_id = locations.id").
		Joins("JOIN places on tickets.place_id = places.id").
		Joins("JOIN sectors on places.sector_id = sectors.id").
		Joins("JOIN prices on performances.id = prices.performance_id and sectors.id = prices.sector_id").
		Rows()
	if err != nil {
		return nil, fmt.Errorf("can't read users from database, error:%w", err)
	}
	for rows.Next() {
		temp := SelectTicket{}
		err := rows.Scan(&temp.Id, &temp.PerformanceName, &temp.GenreName, &temp.PerformanceDuration,
			&temp.DateTime, &temp.HallName, &temp.HallCapacity, &temp.LocationAddress,
			&temp.LocationPhoneNumber, &temp.SectorName, &temp.Place, &temp.Price, &temp.DateOfIssue,
			&temp.Paid, &temp.Reservation, &temp.Destroyed)
		if err != nil {
			return nil, fmt.Errorf("can't scan tickets from database, error:%w", err)
		}
		tickets = append(tickets, temp)
	}
	return tickets, nil
}
