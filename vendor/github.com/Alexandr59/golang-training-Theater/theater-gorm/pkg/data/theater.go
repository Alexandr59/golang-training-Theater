package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Account struct {
	Id          int    `gorm:"primaryKey"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	PhoneNumber string `gorm:"phone_number"`
	Email       string `gorm:"email"`
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

type SelectUser struct {
	Id                  int
	FirstName           string
	LastName            string
	Role                string
	LocationAddress     string
	LocationPhoneNumber string
	PhoneNumber         string
}

type User struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	RoleId      int    `gorm:"role_id"`
	LocationId  int    `gorm:"location_id"`
	PhoneNumber string `gorm:"phone_number"`
}

type Hall struct {
	Id         int    `gorm:"primaryKey"`
	AccountId  int    `gorm:"account_id"`
	Name       string `gorm:"name"`
	Capacity   int    `gorm:"capacity"`
	LocationId int    `gorm:"location_id"`
}

type Location struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	Address     string `gorm:"address"`
	PhoneNumber string `gorm:"phone_number"`
}

type Performance struct {
	Id        int    `gorm:"primaryKey"`
	AccountId int    `gorm:"account_id"`
	Name      string `gorm:"name"`
	GenreId   int    `gorm:"genre_id"`
	Duration  string `gorm:"duration"`
}

type Place struct {
	Id       int    `gorm:"primaryKey"`
	SectorId int    `gorm:"sector_id"`
	Name     string `gorm:"name"`
}

type Poster struct {
	Id         int    `gorm:"primaryKey"`
	AccountId  int    `gorm:"account_id"`
	ScheduleId int    `gorm:"schedule_id"`
	Comment    string `gorm:"comment"`
}

type Price struct {
	Id            int `gorm:"primaryKey"`
	AccountId     int `gorm:"account_id"`
	SectorId      int `gorm:"sector_id"`
	PerformanceId int `gorm:"performance_id"`
	Price         int `gorm:"price"`
}

type Role struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type Schedule struct {
	Id            int    `gorm:"primaryKey"`
	AccountId     int    `gorm:"account_id"`
	PerformanceId int    `gorm:"performance_id"`
	Date          string `gorm:"date"`
	HallId        int    `gorm:"hall_id"`
}

type Sector struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type Genre struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type TheaterData struct {
	db *gorm.DB
}

func NewTheaterData(db *gorm.DB) *TheaterData {
	return &TheaterData{db: db}
}

func (u TheaterData) ReadAllTickets() ([]SelectTicket, error) {
	var tickets []SelectTicket
	rows, err := u.db.Table("tickets").Select("tickets.id, performances.name, genres.name, " +
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

func (u TheaterData) ReadAllPosters() ([]SelectPoster, error) {
	var posters []SelectPoster
	rows, err := u.db.Table("posters").Select("posters.id, performances.name, genres.name, " +
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

func (u TheaterData) ReadAllUsers(account Account) ([]SelectUser, error) {
	var users []SelectUser
	rows, err := u.db.Table("users").Select("users.id, users.first_name, "+
		"users.last_name, roles.name, locations.address, locations.phone_number, users.phone_number").
		Joins("JOIN roles on users.role_id = roles.id").
		Joins("JOIN locations on locations.id = users.account_id").
		Where("users.account_id = ?", account.Id).
		Rows()
	if err != nil {
		return nil, fmt.Errorf("can't get users from database, error:%w", err)
	}
	for rows.Next() {
		var temp SelectUser
		err = rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Role,
			&temp.LocationAddress, &temp.LocationPhoneNumber, &temp.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("can't scan users from database, error:%w", err)
		}
		users = append(users, temp)
	}
	return users, nil
}

func (u TheaterData) AddAccount(account Account) (int, error) {
	result := u.db.Create(&account)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser account to database, error: %w", result.Error)
	}
	return account.Id, nil
}

func (u TheaterData) AddGenre(genre Genre) (int, error) {
	result := u.db.Create(&genre)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser genre to database, error: %w", result.Error)
	}
	return genre.Id, nil
}

func (u TheaterData) AddHall(hall Hall) (int, error) {
	result := u.db.Create(&hall)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser hall to database, error: %w", result.Error)
	}
	return hall.Id, nil
}

func (u TheaterData) AddLocation(location Location) (int, error) {
	result := u.db.Create(&location)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser location to database, error: %w", result.Error)
	}
	return location.Id, nil
}

func (u TheaterData) AddPerformance(performance Performance) (int, error) {
	result := u.db.Create(&performance)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Performance to database, error: %w", result.Error)
	}
	return performance.Id, nil
}

func (u TheaterData) AddPlace(place Place) (int, error) {
	result := u.db.Create(&place)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Place to database, error: %w", result.Error)
	}
	return place.Id, nil
}

func (u TheaterData) AddPoster(poster Poster) (int, error) {
	result := u.db.Create(&poster)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Poster to database, error: %w", result.Error)
	}
	return poster.Id, nil
}

func (u TheaterData) AddPrice(price Price) (int, error) {
	result := u.db.Create(&price)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Price to database, error: %w", result.Error)
	}
	return price.Id, nil
}

func (u TheaterData) AddRole(role Role) (int, error) {
	result := u.db.Create(&role)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Role to database, error: %w", result.Error)
	}
	return role.Id, nil
}

func (u TheaterData) AddSchedule(schedule Schedule) (int, error) {
	result := u.db.Create(&schedule)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Schedule to database, error: %w", result.Error)
	}
	return schedule.Id, nil
}

func (u TheaterData) AddSector(sector Sector) (int, error) {
	result := u.db.Create(&sector)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Sector to database, error: %w", result.Error)
	}
	return sector.Id, nil
}

func (u TheaterData) AddTicket(ticket Ticket) (int, error) {
	result := u.db.Create(&ticket)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser Ticket to database, error: %w", result.Error)
	}
	return ticket.Id, nil
}

func (u TheaterData) AddUser(user User) (int, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return -1, fmt.Errorf("can't inser User to database, error: %w", result.Error)
	}
	return user.Id, nil
}

func (u TheaterData) DeleteAccount(entry Account) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Account to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteGenre(entry Genre) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Genre to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteHall(entry Hall) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Hall to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteLocation(entry Location) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Location to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePerformance(entry Performance) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Performance to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePlace(entry Place) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Place to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePoster(entry Poster) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Poster to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePrice(entry Price) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Price to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteRole(entry Role) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Role to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteSchedule(entry Schedule) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Schedule to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteSector(entry Sector) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Sector to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteTicket(entry Ticket) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Ticket to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteUser(entry User) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete User to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateAccount(account Account) error {
	result := u.db.Model(&account).Updates(account)
	if result.Error != nil {
		return fmt.Errorf("can't update account to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateGenre(genre Genre) error {
	result := u.db.Model(&genre).Updates(genre)
	if result.Error != nil {
		return fmt.Errorf("can't update genre to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateHall(hall Hall) error {
	result := u.db.Model(&hall).Updates(hall)
	if result.Error != nil {
		return fmt.Errorf("can't update hall to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateLocation(location Location) error {
	result := u.db.Model(&location).Updates(location)
	if result.Error != nil {
		return fmt.Errorf("can't update location to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePerformance(performance Performance) error {
	result := u.db.Model(&performance).Updates(performance)
	if result.Error != nil {
		return fmt.Errorf("can't update Performance to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePlace(place Place) error {
	result := u.db.Model(&place).Updates(place)
	if result.Error != nil {
		return fmt.Errorf("can't update Place to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePoster(poster Poster) error {
	result := u.db.Model(&poster).Updates(poster)
	if result.Error != nil {
		return fmt.Errorf("can't update Poster to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePrice(price Price) error {
	result := u.db.Model(&price).Updates(price)
	if result.Error != nil {
		return fmt.Errorf("can't update Price to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateRole(role Role) error {
	result := u.db.Model(&role).Updates(role)
	if result.Error != nil {
		return fmt.Errorf("can't update Role to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateSchedule(schedule Schedule) error {
	result := u.db.Model(&schedule).Updates(schedule)
	if result.Error != nil {
		return fmt.Errorf("can't update Schedule to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateSector(sector Sector) error {
	result := u.db.Model(&sector).Updates(sector)
	if result.Error != nil {
		return fmt.Errorf("can't update Sector to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateTicket(ticket Ticket) error {
	result := u.db.Model(&ticket).Updates(ticket)
	if result.Error != nil {
		return fmt.Errorf("can't update Ticket to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateUser(user User) error {
	result := u.db.Model(&user).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("can't update User to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) FindByIdAccount(entry Account) (Account, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Account{}, fmt.Errorf("can't find Account to database, error: %w", result.Error)
	}
	return entry, nil
}
func (u TheaterData) FindByIdGenre(entry Genre) (Genre, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Genre{}, fmt.Errorf("can't find Genre to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdHall(entry Hall) (Hall, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Hall{}, fmt.Errorf("can't find Hall to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdLocation(entry Location) (Location, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Location{}, fmt.Errorf("can't find Location to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdPerformance(entry Performance) (Performance, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Performance{}, fmt.Errorf("can't find Performance to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdPlace(entry Place) (Place, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Place{}, fmt.Errorf("can't find Place to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdPoster(entry Poster) (Poster, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Poster{}, fmt.Errorf("can't find Poster to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdPrice(entry Price) (Price, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Price{}, fmt.Errorf("can't find Price to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdRole(entry Role) (Role, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Role{}, fmt.Errorf("can't find Role to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdSchedule(entry Schedule) (Schedule, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Schedule{}, fmt.Errorf("can't find Schedule to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdSector(entry Sector) (Sector, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Sector{}, fmt.Errorf("can't find Sector to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdTicket(entry Ticket) (Ticket, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return Ticket{}, fmt.Errorf("can't find Ticket to database, error: %w", result.Error)
	}
	return entry, nil
}

func (u TheaterData) FindByIdUser(entry User) (User, error) {
	result := u.db.First(&entry)
	if result.Error != nil {
		return User{}, fmt.Errorf("can't find User to database, error: %w", result.Error)
	}
	return entry, nil
}
