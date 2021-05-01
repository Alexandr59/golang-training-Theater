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

type TicketSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *TicketData
}

func TestInitTicket(t *testing.T) {
	suite.Run(t, new(TicketSuite))
}

func (s *TicketSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *TicketSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewTicketData(s.DB)
}

var testTicket = &Ticket{
	Id:          23,
	AccountId:   1,
	ScheduleId:  10,
	PlaceId:     10,
	DateOfIssue: "now()",
	Paid:        true,
	Reservation: true,
	Destroyed:   false,
}

var testSelectTicket = &SelectTicket{
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

func (s *TicketSuite) TestTheaterData_AddTicket() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tickets"`)).
		WithArgs(testTicket.Id, testTicket.AccountId, testTicket.ScheduleId,
			testTicket.PlaceId, testTicket.DateOfIssue, testTicket.Paid,
			testTicket.Reservation, testTicket.Destroyed).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddTicket(*testTicket)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *TicketSuite) TestTheaterData_AddTicketErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tickets"`)).
		WithArgs(testTicket.AccountId, testTicket.ScheduleId,
			testTicket.PlaceId, testTicket.DateOfIssue, testTicket.Paid,
			testTicket.Reservation, testTicket.Destroyed).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddTicket(*testTicket)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *TicketSuite) TestTheaterData_DeleteTicket() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tickets"`)).
		WithArgs(testTicket.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteTicket(*testTicket)
	require.NoError(s.T(), err)
}

func (s *TicketSuite) TestTheaterData_DeleteTicketErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tickets"`)).
		WithArgs(testTicket.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteTicket(*testTicket)
	require.Error(s.T(), err)
}

func (s *TicketSuite) TestTheaterData_UpdateTicket() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tickets"`)).
		WithArgs(testTicket.AccountId, testTicket.DateOfIssue,
			testTicket.Id, testTicket.Paid, testTicket.PlaceId,
			testTicket.Reservation, testTicket.ScheduleId, testTicket.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateTicket(*testTicket)
	require.NoError(s.T(), err)
}

func (s *TicketSuite) TestTheaterData_UpdateTicketErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tickets"`)).
		WithArgs(testTicket.AccountId, testTicket.DateOfIssue,
			testTicket.Id, testTicket.Paid, testTicket.PlaceId,
			testTicket.Reservation, testTicket.ScheduleId, testTicket.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateTicket(*testTicket)
	require.Error(s.T(), err)
}

func (s *TicketSuite) TestTheaterData_FindByIdTicket() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "schedule_id",
		"place_id", "date_of_issue", "paid", "reservation", "destroyed"}).
		AddRow(testTicket.Id, testTicket.AccountId, testTicket.ScheduleId,
			testTicket.PlaceId, testTicket.DateOfIssue, testTicket.Paid, testTicket.Reservation, testTicket.Destroyed)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tickets" WHERE "tickets"."id" = $1 ORDER BY "tickets"."id" ASC LIMIT 1`)).
		WithArgs(testTicket.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdTicket(Ticket{Id: 23})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testTicket, &res))
}

func (s *TicketSuite) TestTheaterData_FindByIdTicketErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tickets" WHERE "tickets"."id" = $1 ORDER BY "tickets"."id" ASC LIMIT 1`)).
		WithArgs(testTicket.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdTicket(Ticket{Id: 23})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *TicketSuite) TestTheaterData_ReadAllTickets() {
	rows := sqlmock.NewRows([]string{"tickets.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "sectors.name", "places.name", "price.price",
		"tickets.date_of_issue", "tickets.paid", "tickets.reservation", "tickets.destroyed"}).
		AddRow(testSelectTicket.Id, testSelectTicket.PerformanceName, testSelectTicket.GenreName, testSelectTicket.PerformanceDuration,
			testSelectTicket.DateTime, testSelectTicket.HallName, testSelectTicket.HallCapacity, testSelectTicket.LocationAddress, testSelectTicket.LocationPhoneNumber,
			testSelectTicket.SectorName, testSelectTicket.Place, testSelectTicket.Price, testSelectTicket.DateOfIssue, testSelectTicket.Paid, testSelectTicket.Reservation,
			testSelectTicket.Destroyed)
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
	require.Nil(s.T(), deep.Equal(testSelectTicket, &res[0]))
}

func (s *TicketSuite) TestTheaterData_ReadAllTicketsErr() {
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
