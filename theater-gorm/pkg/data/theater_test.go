package data

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *TheaterData
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewTheaterData(s.DB)
}

var testTicket = &SelectTicket{
	Id:                  21,
	PerformanceName:     "The Dragon",
	GenreName:           "a musical",
	PerformanceDuration: "0000-01-01T04:00:00Z",
	DateTime:            "2021-04-13T16:00:00Z",
	HallName:            "Middle",
	HallCapacity:        1500,
	LocationAddress:     "Gaidara_6",
	LocationPhoneNumber: "+375443564987",
	SectorName:          "A",
	Place:               1,
	Price:               40,
	DateOfIssue:         "2021-04-12T22:48:15.344148Z",
	Paid:                false,
	Reservation:         false,
	Destroyed:           false,
}

var testPoster = &SelectPoster{
	Id:                  2,
	PerformanceName:     "The Dragon",
	GenreName:           "a musical",
	PerformanceDuration: "0000-01-01T04:00:00Z",
	DateTime:            "2021-04-13T16:00:00Z",
	HallName:            "Middle",
	HallCapacity:        1500,
	LocationAddress:     "Gaidara_6",
	LocationPhoneNumber: "+375443564987",
	Comment:             "We invite you! It will be cool!!!",
}

var testUser = &SelectUser{
	Id:                  1,
	FirstName:           "Charles",
	LastName:            "Dean",
	Role:                "Actor",
	LocationAddress:     "Gaidara_6",
	LocationPhoneNumber: "+375443564987",
	PhoneNumber:         "+375445239375",
}

var testAccount = &Account{
	Id:          4,
	FirstName:   "Dim",
	LastName:    "Ivanov",
	PhoneNumber: "+375296574897",
	Email:       "dimaivanov@gmail.com",
}

var testGenre = &Genre{
	Id:   1,
	Name: "a musical",
}

var testHall = &Hall{
	Id:         4,
	AccountId:  1,
	Name:       "testName",
	Capacity:   1099,
	LocationId: 1,
}

var testLocation = &Location{
	Id:          4,
	AccountId:   1,
	Address:     "Gaidara10",
	PhoneNumber: "+3754466633321",
}

var testPerformance = &Performance{
	Id:        4,
	AccountId: 1,
	Name:      "Big ball",
	GenreId:   3,
	Duration:  "1:00",
}

var testPlace = &Place{
	Id:       6,
	SectorId: 15,
	Name:     "2",
}

var testPoster1 = &Poster{
	Id:         2,
	AccountId:  1,
	ScheduleId: 1,
	Comment:    "Hi!!!",
}

var testPrice = &Price{
	Id:            6,
	AccountId:     1,
	SectorId:      10,
	PerformanceId: 2,
	Price:         120,
}

var testRole = &Role{
	Id:   6,
	Name: "Test",
}

var testSchedule = &Schedule{
	Id:            8,
	AccountId:     1,
	PerformanceId: 3,
	Date:          "2021-04-13 16:00",
	HallId:        3,
}

var testSector = &Sector{
	Id:   9,
	Name: "L",
}

var testTicket1 = &Ticket{
	Id:          23,
	AccountId:   1,
	ScheduleId:  10,
	PlaceId:     10,
	DateOfIssue: "now()",
	Paid:        true,
	Reservation: true,
	Destroyed:   false,
}

var testUser1 = &User{
	Id:          3,
	AccountId:   1,
	FirstName:   "TestFirstName",
	LastName:    "TestLastName",
	RoleId:      3,
	LocationId:  1,
	PhoneNumber: "+3753347362873267",
}

func (s *Suite) TestTheaterData_ReadAllTickets() {
	rows := sqlmock.NewRows([]string{"tickets.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "sectors.name", "places.name", "price.price",
		"tickets.date_of_issue", "tickets.paid", "tickets.reservation", "tickets.destroyed"}).
		AddRow(testTicket.Id, testTicket.PerformanceName, testTicket.GenreName, testTicket.PerformanceDuration,
			testTicket.DateTime, testTicket.HallName, testTicket.HallCapacity, testTicket.LocationAddress, testTicket.LocationPhoneNumber,
			testTicket.SectorName, testTicket.Place, testTicket.Price, testTicket.DateOfIssue, testTicket.Paid, testTicket.Reservation,
			testTicket.Destroyed)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT tickets.id, performances.name, genres.name, performances.duration, schedules.date, halls.name,
halls.capacity, locations.address, locations.phone_number, sectors.name, places.name, prices.price, tickets.date_of_issue, tickets.paid, 
tickets.reservation, tickets.destroyed FROM "tickets" 
JOIN schedules on schedules.id = tickets.schedule_id 
JOIN performances on schedules.performance_id = performances.id 
JOIN genres on performances.genre_id = genres.id 
JOIN halls on schedules.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id 
JOIN places on tickets.place_id = places.id JOIN sectors on places.sector_id = sectors.id 
JOIN prices on performances.id = prices.performance_id and sectors.id = prices.sector_id`)).
		WillReturnRows(rows)
	res, err := s.data.ReadAllTickets()
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testTicket, &res[0]))
}

func (s *Suite) TestTheaterData_ReadAllTicketsErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT tickets.id, performances.name, genres.name, performances.duration, schedules.date, halls.name, 
halls.capacity, locations.address, locations.phone_number, sectors.name, places.name, prices.price, tickets.date_of_issue, tickets.paid, 
tickets.reservation, tickets.destroyed FROM "tickets" 
JOIN schedules on schedules.id = tickets.schedule_id 
JOIN performances on schedules.performance_id = performances.id 
JOIN genres on performances.genre_id = genres.id 
JOIN halls on schedules.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id 
JOIN places on tickets.place_id = places.id JOIN sectors on places.sector_id = sectors.id 
JOIN prices on performances.id = prices.performance_id and sectors.id = prices.sector_id`)).
		WillReturnError(errors.New("something went wrong"))
	users, err := s.data.ReadAllTickets()
	require.Error(s.T(), err)
	require.Empty(s.T(), users)
}

func (s *Suite) TestTheaterData_ReadAllPosters() {
	rows := sqlmock.NewRows([]string{"poster.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "poster.comment"}).
		AddRow(testPoster.Id, testPoster.PerformanceName, testPoster.GenreName, testPoster.PerformanceDuration,
			testPoster.DateTime, testPoster.HallName, testPoster.HallCapacity, testPoster.LocationAddress,
			testPoster.LocationPhoneNumber, testPoster.Comment)
	s.mock.ExpectQuery(`SELECT posters.id, performances.name, genres.name, performances.duration, schedules.date, 
halls.name, halls.capacity, locations.address, locations.phone_number, posters.comment FROM "posters" 
JOIN schedules on schedules.id = posters.schedule_id 
JOIN performances on schedules.performance_id = performances.id 
JOIN genres on performances.genre_id = genres.id 
JOIN halls on schedules.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id`).
		WillReturnRows(rows)
	res, err := s.data.ReadAllPosters()
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPoster, &res[0]))
}

func (s *Suite) TestTheaterData_ReadAllPostersErr() {
	s.mock.ExpectQuery(`SELECT posters.id, performances.name, genres.name, performances.duration, schedules.date, 
halls.name, halls.capacity, locations.address, locations.phone_number, posters.comment FROM "posters" 
JOIN schedules on schedules.id = posters.schedule_id 
JOIN performances on schedules.performance_id = performances.id 
JOIN genres on performances.genre_id = genres.id 
JOIN halls on schedules.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id`).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.ReadAllPosters()
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_ReadAllUsers() {
	rows := sqlmock.NewRows([]string{"users.id", "users.first_name", "users.last_name", "roles.name", "locations.address",
		"locations.phone_number", "users.phone_number"}).
		AddRow(testUser.Id, testUser.FirstName, testUser.LastName, testUser.Role,
			testUser.LocationAddress, testUser.LocationPhoneNumber, testUser.PhoneNumber)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.first_name, users.last_name, roles.name, locations.address, 
locations.phone_number, users.phone_number FROM "users" 
JOIN roles on users.role_id = roles.id 
JOIN locations on locations.id = users.account_id 
WHERE (users.account_id = $1)`)).
		WithArgs(Account{Id: 1}.Id).
		WillReturnRows(rows)
	res, err := s.data.ReadAllUsers(Account{Id: 1})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testUser, &res[0]))
}

func (s *Suite) TestTheaterData_ReadAllUsersErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.first_name, users.last_name, roles.name, locations.address, 
locations.phone_number, users.phone_number FROM "users" 
JOIN roles on users.role_id = roles.id 
JOIN locations on locations.id = users.account_id 
WHERE (users.account_id = $1)`)).
		WithArgs(Account{Id: 1}.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.ReadAllUsers(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdAccount() {
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "phone_number", "email"}).
		AddRow(testAccount.Id, testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdAccount(Account{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testAccount, &res))
}

func (s *Suite) TestTheaterData_FindByIdAccountErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdAccount(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdGenre() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(testGenre.Id, testGenre.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "genres" WHERE "genres"."id" = $1 ORDER BY "genres"."id" ASC LIMIT 1`)).
		WithArgs(testGenre.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdGenre(Genre{Id: 1})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testGenre, &res))
}

func (s *Suite) TestTheaterData_FindByIdGenreErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "genres" WHERE "genres"."id" = $1 ORDER BY "genres"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdAccount(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdHall() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "name", "capacity", "location_id"}).
		AddRow(testHall.Id, testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "halls" WHERE "halls"."id" = $1 ORDER BY "halls"."id" ASC LIMIT 1`)).
		WithArgs(testHall.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdHall(Hall{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testHall, &res))
}

func (s *Suite) TestTheaterData_FindByIdHallErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "halls" WHERE "halls"."id" = $1 ORDER BY "halls"."id" ASC LIMIT 1`)).
		WithArgs(testHall.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdHall(Hall{Id: 4})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdLocation() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "address", "phone_number"}).
		AddRow(testLocation.Id, testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "locations" WHERE "locations"."id" = $1 ORDER BY "locations"."id" ASC LIMIT 1`)).
		WithArgs(testLocation.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdLocation(Location{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testLocation, &res))
}

func (s *Suite) TestTheaterData_FindByIdLocationErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "locations" WHERE "locations"."id" = $1 ORDER BY "locations"."id" ASC LIMIT 1`)).
		WithArgs(testLocation.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdLocation(Location{Id: 4})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdPerformance() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "name", "genre_id", "duration"}).
		AddRow(testPerformance.Id, testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "performances" WHERE "performances"."id" = $1 ORDER BY "performances"."id" ASC LIMIT 1`)).
		WithArgs(testPerformance.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPerformance(Performance{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPerformance, &res))
}

func (s *Suite) TestTheaterData_FindByIdPerformanceErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "performances" WHERE "performances"."id" = $1 ORDER BY "performances"."id" ASC LIMIT 1`)).
		WithArgs(testPerformance.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPerformance(Performance{Id: 4})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdPlace() {
	rows := sqlmock.NewRows([]string{"id", "sector_id", "name"}).
		AddRow(testPlace.Id, testPlace.SectorId, testPlace.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "places" WHERE "places"."id" = $1 ORDER BY "places"."id" ASC LIMIT 1`)).
		WithArgs(testPlace.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPlace(Place{Id: 6})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPlace, &res))
}

func (s *Suite) TestTheaterData_FindByIdPlaceErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "places" WHERE "places"."id" = $1 ORDER BY "places"."id" ASC LIMIT 1`)).
		WithArgs(testPlace.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPlace(Place{Id: 6})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdPoster() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "schedule_id", "comment"}).
		AddRow(testPoster1.Id, testPoster1.AccountId, testPoster1.ScheduleId, testPoster1.Comment)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posters" WHERE "posters"."id" = $1 ORDER BY "posters"."id" ASC LIMIT 1`)).
		WithArgs(testPoster1.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPoster(Poster{Id: 2})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPoster1, &res))
}

func (s *Suite) TestTheaterData_FindByIdPosterErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posters" WHERE "posters"."id" = $1 ORDER BY "posters"."id" ASC LIMIT 1`)).
		WithArgs(testPoster1.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPoster(Poster{Id: 2})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdPrice() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "sector_id", "performance_id", "price"}).
		AddRow(testPrice.Id, testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "prices" WHERE "prices"."id" = $1 ORDER BY "prices"."id" ASC LIMIT 1`)).
		WithArgs(testPrice.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPrice(Price{Id: 6})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPrice, &res))
}

func (s *Suite) TestTheaterData_FindByIdPriceErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "prices" WHERE "prices"."id" = $1 ORDER BY "prices"."id" ASC LIMIT 1`)).
		WithArgs(testPrice.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPrice(Price{Id: 6})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdRole() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(testRole.Id, testRole.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "roles" WHERE "roles"."id" = $1 ORDER BY "roles"."id" ASC LIMIT 1`)).
		WithArgs(testRole.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdRole(Role{Id: 6})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testRole, &res))
}

func (s *Suite) TestTheaterData_FindByIdRoleErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "roles" WHERE "roles"."id" = $1 ORDER BY "roles"."id" ASC LIMIT 1`)).
		WithArgs(testRole.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdRole(Role{Id: 6})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdSchedule() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "performance_id", "date", "hall_id"}).
		AddRow(testSchedule.Id, testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "schedules" WHERE "schedules"."id" = $1 ORDER BY "schedules"."id" ASC LIMIT 1`)).
		WithArgs(testSchedule.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdSchedule(Schedule{Id: 8})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testSchedule, &res))
}

func (s *Suite) TestTheaterData_FindByIdScheduleErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "schedules" WHERE "schedules"."id" = $1 ORDER BY "schedules"."id" ASC LIMIT 1`)).
		WithArgs(testSchedule.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdSchedule(Schedule{Id: 8})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdSector() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(testSector.Id, testSector.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "sectors" WHERE "sectors"."id" = $1 ORDER BY "sectors"."id" ASC LIMIT 1`)).
		WithArgs(testSector.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdSector(Sector{Id: 9})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testSector, &res))
}

func (s *Suite) TestTheaterData_FindByIdSectorErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "sectors" WHERE "sectors"."id" = $1 ORDER BY "sectors"."id" ASC LIMIT 1`)).
		WithArgs(testSector.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdSector(Sector{Id: 9})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdTicket() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "schedule_id",
		"place_id", "date_of_issue", "paid", "reservation", "destroyed"}).
		AddRow(testTicket1.Id, testTicket1.AccountId, testTicket1.ScheduleId,
			testTicket1.PlaceId, testTicket1.DateOfIssue, testTicket1.Paid, testTicket1.Reservation, testTicket1.Destroyed)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tickets" WHERE "tickets"."id" = $1 ORDER BY "tickets"."id" ASC LIMIT 1`)).
		WithArgs(testTicket1.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdTicket(Ticket{Id: 23})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testTicket1, &res))
}

func (s *Suite) TestTheaterData_FindByIdTicketErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tickets" WHERE "tickets"."id" = $1 ORDER BY "tickets"."id" ASC LIMIT 1`)).
		WithArgs(testTicket1.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdTicket(Ticket{Id: 23})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdUser() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "first_name", "last_name",
		"role_id", "location_id", "phone_number"}).
		AddRow(testUser1.Id, testUser1.AccountId, testUser1.FirstName,
			testUser1.LastName, testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(testUser1.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdUser(User{Id: 3})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testUser1, &res))
}

func (s *Suite) TestTheaterData_FindByIdUserErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(testUser1.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdUser(User{Id: 3})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_AddAccount() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts"`)).
		WithArgs(testAccount.Id, testAccount.FirstName, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddAccount(*testAccount)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddAccountErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts"`)).
		WithArgs(testAccount.Id, testAccount.FirstName, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Email).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddAccount(*testAccount)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddGenre() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "genres"`)).
		WithArgs(testGenre.Id, testGenre.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddGenre(*testGenre)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddGenreErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "genres" ("id","name") VALUES ($1,$2) RETURNING "genres"."id"`)).
		WithArgs(testGenre.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddGenre(*testGenre)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddHall() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "halls"`)).
		WithArgs(testHall.Id, testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddHall(*testHall)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddHallErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "halls"`)).
		WithArgs(testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddHall(*testHall)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddLocation() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "locations"`)).
		WithArgs(testLocation.Id, testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddLocation(*testLocation)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddLocationErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "locations"`)).
		WithArgs(testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddLocation(*testLocation)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddPerformance() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "performances"`)).
		WithArgs(testPerformance.Id, testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPerformance(*testPerformance)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddPerformanceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "performances"`)).
		WithArgs(testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPerformance(*testPerformance)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddPlace() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "places"`)).
		WithArgs(testPlace.Id, testPlace.SectorId, testPlace.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPlace(*testPlace)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddPlaceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "places"`)).
		WithArgs(testPlace.SectorId, testPlace.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPlace(*testPlace)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddPoster() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "posters"`)).
		WithArgs(testPoster1.Id, testPoster1.AccountId, testPoster1.ScheduleId, testPoster1.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPoster(*testPoster1)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddPosterErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "posters"`)).
		WithArgs(testPoster1.AccountId, testPoster1.ScheduleId, testPoster1.Comment).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPoster(*testPoster1)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddPrice() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "prices"`)).
		WithArgs(testPrice.Id, testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPrice(*testPrice)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}
func (s *Suite) TestTheaterData_AddPriceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "prices"`)).
		WithArgs(testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPrice(*testPrice)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddRole() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "roles"`)).
		WithArgs(testRole.Id, testRole.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddRole(*testRole)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddRoleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "roles"`)).
		WithArgs(testRole.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddRole(*testRole)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddSchedule() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "schedules"`)).
		WithArgs(testSchedule.Id, testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddSchedule(*testSchedule)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddScheduleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "schedules"`)).
		WithArgs(testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddSchedule(*testSchedule)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddSector() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "sectors"`)).
		WithArgs(testSector.Id, testSector.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddSector(*testSector)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddSectorErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "sectors"`)).
		WithArgs(testSector.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddSector(*testSector)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddTicket() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tickets"`)).
		WithArgs(testTicket1.Id, testTicket1.AccountId, testTicket1.ScheduleId,
			testTicket1.PlaceId, testTicket1.DateOfIssue, testTicket1.Paid,
			testTicket1.Reservation, testTicket1.Destroyed).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddTicket(*testTicket1)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddTicketErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tickets"`)).
		WithArgs(testTicket1.AccountId, testTicket1.ScheduleId,
			testTicket1.PlaceId, testTicket1.DateOfIssue, testTicket1.Paid,
			testTicket1.Reservation, testTicket1.Destroyed).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddTicket(*testTicket1)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_AddUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs(testUser1.Id, testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddUser(*testUser1)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *Suite) TestTheaterData_AddUserErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddUser(*testUser1)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *Suite) TestTheaterData_UpdateAccount() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "accounts"`)).
		WithArgs(testAccount.Email, testAccount.FirstName, testAccount.Id, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateAccount(*testAccount)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateAccountErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "accounts"`)).
		WithArgs(testAccount.Email, testAccount.FirstName, testAccount.Id, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateAccount(*testAccount)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateGenre() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "genres"`)).
		WithArgs(testGenre.Id, testGenre.Name, testGenre.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateGenre(*testGenre)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateGenreErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "genres"`)).
		WithArgs(testGenre.Id, testGenre.Name, testGenre.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateGenre(*testGenre)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateHall() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "halls"`)).
		WithArgs(testHall.AccountId, testHall.Capacity, testHall.Id,
			testHall.LocationId, testHall.Name, testHall.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateHall(*testHall)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateHallErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "halls"`)).
		WithArgs(testHall.AccountId, testHall.Capacity, testHall.Id,
			testHall.LocationId, testHall.Name, testHall.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateHall(*testHall)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateLocation() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "locations"`)).
		WithArgs(testLocation.AccountId, testLocation.Address,
			testLocation.Id, testLocation.PhoneNumber, testLocation.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateLocation(*testLocation)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateLocationErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "locations"`)).
		WithArgs(testLocation.AccountId, testLocation.Address,
			testLocation.Id, testLocation.PhoneNumber, testLocation.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateLocation(*testLocation)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePerformance() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "performances"`)).
		WithArgs(testPerformance.AccountId, testPerformance.Duration,
			testPerformance.GenreId, testPerformance.Id, testPerformance.Name, testPerformance.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePerformance(*testPerformance)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePerformanceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "performances"`)).
		WithArgs(testPerformance.AccountId, testPerformance.Duration,
			testPerformance.GenreId, testPerformance.Id, testPerformance.Name, testPerformance.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePerformance(*testPerformance)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePlace() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "places"`)).
		WithArgs(testPlace.Id, testPlace.Name, testPlace.SectorId, testPlace.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePlace(*testPlace)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePlaceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "places"`)).
		WithArgs(testPlace.Id, testPlace.Name, testPlace.SectorId, testPlace.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePlace(*testPlace)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePoster() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posters"`)).
		WithArgs(testPoster1.AccountId, testPoster1.Comment,
			testPoster1.Id, testPoster1.ScheduleId, testPoster1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePoster(*testPoster1)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePosterErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posters"`)).
		WithArgs(testPoster1.AccountId, testPoster1.Comment,
			testPoster1.Id, testPoster1.ScheduleId, testPoster1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePoster(*testPoster1)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePrice() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "prices"`)).
		WithArgs(testPrice.AccountId, testPrice.Id, testPrice.PerformanceId,
			testPrice.Price, testPrice.SectorId, testPrice.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePrice(*testPrice)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdatePriceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "prices"`)).
		WithArgs(testPrice.AccountId, testPrice.Id, testPrice.PerformanceId,
			testPrice.Price, testPrice.SectorId, testPrice.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePrice(*testPrice)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateRole() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "roles"`)).
		WithArgs(testRole.Id, testRole.Name, testRole.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateRole(*testRole)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateRoleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "roles"`)).
		WithArgs(testRole.Id, testRole.Name, testRole.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateRole(*testRole)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateSchedule() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "schedules"`)).
		WithArgs(testSchedule.AccountId, testSchedule.Date, testSchedule.HallId,
			testSchedule.Id, testSchedule.PerformanceId, testSchedule.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateSchedule(*testSchedule)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateScheduleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "schedules"`)).
		WithArgs(testSchedule.AccountId, testSchedule.Date, testSchedule.HallId,
			testSchedule.Id, testSchedule.PerformanceId, testSchedule.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateSchedule(*testSchedule)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateSector() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "sectors"`)).
		WithArgs(testSector.Id, testSector.Name, testSector.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateSector(*testSector)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateSectorErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "sectors"`)).
		WithArgs(testSector.Id, testSector.Name, testSector.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateSector(*testSector)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateTicket() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tickets"`)).
		WithArgs(testTicket1.AccountId, testTicket1.DateOfIssue,
			testTicket1.Id, testTicket1.Paid, testTicket1.PlaceId,
			testTicket1.Reservation, testTicket1.ScheduleId, testTicket1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateTicket(*testTicket1)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateTicketErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tickets"`)).
		WithArgs(testTicket1.AccountId, testTicket1.DateOfIssue,
			testTicket1.Id, testTicket1.Paid, testTicket1.PlaceId,
			testTicket1.Reservation, testTicket1.ScheduleId, testTicket1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateTicket(*testTicket1)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.Id, testUser1.LastName,
			testUser1.LocationId, testUser1.PhoneNumber, testUser1.RoleId, testUser1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateUser(*testUser1)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_UpdateUserErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.Id, testUser1.LastName,
			testUser1.LocationId, testUser1.PhoneNumber, testUser1.RoleId, testUser1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateUser(*testUser1)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteAccount() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "accounts"`)).
		WithArgs(testAccount.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteAccount(*testAccount)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteAccountErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "accounts"`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteAccount(*testAccount)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteGenre() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "genres"`)).
		WithArgs(testGenre.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteGenre(*testGenre)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteGenreErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "genres"`)).
		WithArgs(testGenre.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteGenre(*testGenre)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteHall() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "halls"`)).
		WithArgs(testHall.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteHall(*testHall)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteHallErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "halls"`)).
		WithArgs(testHall.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteHall(*testHall)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteLocation() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "locations"`)).
		WithArgs(testLocation.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteLocation(*testLocation)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteLocationErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "locations"`)).
		WithArgs(testLocation.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteLocation(*testLocation)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePerformance() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "performances"`)).
		WithArgs(testPerformance.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePerformance(*testPerformance)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePerformanceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "performances"`)).
		WithArgs(testPerformance.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePerformance(*testPerformance)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePlace() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "places"`)).
		WithArgs(testPlace.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePlace(*testPlace)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePlaceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "places"`)).
		WithArgs(testPlace.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePlace(*testPlace)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePoster() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "posters"`)).
		WithArgs(testPoster1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePoster(*testPoster1)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePosterErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "posters"`)).
		WithArgs(testPoster1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePoster(*testPoster1)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePrice() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "prices"`)).
		WithArgs(testPrice.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePrice(*testPrice)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeletePriceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "prices"`)).
		WithArgs(testPrice.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePrice(*testPrice)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteRole() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "roles"`)).
		WithArgs(testRole.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteRole(*testRole)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteRoleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "roles"`)).
		WithArgs(testRole.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteRole(*testRole)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteSchedule() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "schedules"`)).
		WithArgs(testSchedule.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteSchedule(*testSchedule)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteScheduleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "schedules"`)).
		WithArgs(testSchedule.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteSchedule(*testSchedule)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteSector() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "sectors"`)).
		WithArgs(testSector.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteSector(*testSector)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteSectorErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "sectors"`)).
		WithArgs(testSector.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteSector(*testSector)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteTicket() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tickets"`)).
		WithArgs(testTicket1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteTicket(*testTicket1)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteTicketErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tickets"`)).
		WithArgs(testTicket1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteTicket(*testTicket1)
	require.Error(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"`)).
		WithArgs(testUser1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteUser(*testUser1)
	require.NoError(s.T(), err)
}

func (s *Suite) TestTheaterData_DeleteUserErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"`)).
		WithArgs(testUser1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteUser(*testUser1)
	require.Error(s.T(), err)
}
