package data

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
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
	Id:          14,
	FirstName:   "Dim",
	LastName:    "Ivanov",
	PhoneNumber: "+375296574897",
	Email:       "dimaivanov@gmail.com",
}

var testGenre = &Genre{
	Id:   6,
	Name: "a musical",
}

var testHall = &Hall{
	Id:         4,
	AccountId:  1,
	Name:       "",
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

var testTable = DeleteAccounts
var id = 11

func TestTheaterData_ReadAllTickets(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	rows := sqlmock.NewRows([]string{"tickets.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "sectors.name", "places.name", "price.price",
		"tickets.date_of_issue", "tickets.paid", "tickets.reservation", "tickets.destroyed"}).
		AddRow(testTicket.Id, testTicket.PerformanceName, testTicket.GenreName, testTicket.PerformanceDuration,
			testTicket.DateTime, testTicket.HallName, testTicket.HallCapacity, testTicket.LocationAddress, testTicket.LocationPhoneNumber,
			testTicket.SectorName, testTicket.Place, testTicket.Price, testTicket.DateOfIssue, testTicket.Paid, testTicket.Reservation,
			testTicket.Destroyed)
	mock.ExpectQuery(readAllTicketsQuery).WillReturnRows(rows)
	tickets, err := data.ReadAllTickets()
	assert.NoError(err)
	assert.NotEmpty(tickets)
	assert.Equal(tickets[0], *testTicket)
	assert.Len(tickets, 1)
}

func TestTheaterData_ReadAllTicketsErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectQuery(readAllTicketsQuery).WillReturnError(errors.New("something went wrong..."))
	tickets, err := data.ReadAllTickets()
	assert.Error(err)
	assert.Empty(tickets)
}

func TestTheaterData_ReadAllPosters(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	rows := sqlmock.NewRows([]string{"poster.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "poster.comment"}).
		AddRow(testPoster.Id, testPoster.PerformanceName, testPoster.GenreName, testPoster.PerformanceDuration,
			testPoster.DateTime, testPoster.HallName, testPoster.HallCapacity, testPoster.LocationAddress,
			testPoster.LocationPhoneNumber, testPoster.Comment)
	mock.ExpectQuery(readAllPostersQuery).WillReturnRows(rows)
	posters, err := data.ReadAllPosters()
	assert.NoError(err)
	assert.NotEmpty(posters)
	assert.Equal(posters[0], *testPoster)
	assert.Len(posters, 1)
}

func TestTheaterData_ReadAllPostersErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectQuery(readAllPostersQuery).WillReturnError(errors.New("something went wrong..."))
	posters, err := data.ReadAllTickets()
	assert.Error(err)
	assert.Empty(posters)
}

func TestTheaterData_ReadAllUsers(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	rows := sqlmock.NewRows([]string{"users.id", "users.first_name", "users.last_name", "r.name", "a.address",
		"a.phone_number", "users.phone_number"}).
		AddRow(testUser.Id, testUser.FirstName, testUser.LastName, testUser.Role,
			testUser.LocationAddress, testUser.LocationPhoneNumber, testUser.PhoneNumber)
	mock.ExpectQuery(regexp.QuoteMeta(readAllUsersQuery)).
		WithArgs(testUser.Id).
		WillReturnRows(rows)
	users, err := data.ReadAllUsers(Account{Id: 1})
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], *testUser)
	assert.Len(users, 1)
}

func TestTheaterData_ReadAllUsersErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectQuery(readAllPostersQuery).
		WillReturnError(errors.New("something went wrong..."))
	users, err := data.ReadAllTickets()
	assert.Error(err)
	assert.Empty(users)
}

func TestTheaterData_AddAccount(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertAccount)).
		WithArgs(testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddAccount(*testAccount)
	assert.NoError(err)
}

func TestTheaterData_AddAccountErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertAccount)).
		WithArgs(testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddAccount(*testAccount)
	assert.Error(err)
}

func TestTheaterData_AddGenre(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertGenre)).
		WithArgs(testGenre.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddGenre(*testGenre)
	assert.NoError(err)
}

func TestTheaterData_AddGenreErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertGenre)).
		WithArgs(testGenre.Name).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddGenre(*testGenre)
	assert.Error(err)
}

func TestTheaterData_AddHall(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertHall)).
		WithArgs(testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddHall(*testHall)
	assert.NoError(err)
}

func TestTheaterData_AddHallErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertHall)).
		WithArgs(testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddHall(*testHall)
	assert.Error(err)
}

func TestTheaterData_AddLocation(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertLocation)).
		WithArgs(testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddLocation(*testLocation)
	assert.NoError(err)
}

func TestTheaterData_AddLocationErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertLocation)).
		WithArgs(testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddLocation(*testLocation)
	assert.Error(err)
}

func TestTheaterData_AddPerformance(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPerformance)).
		WithArgs(testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddPerformance(*testPerformance)
	assert.NoError(err)
}

func TestTheaterData_AddPerformanceErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPerformance)).
		WithArgs(testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddPerformance(*testPerformance)
	assert.Error(err)
}

func TestTheaterData_AddPlace(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPlace)).
		WithArgs(testPlace.SectorId, testPlace.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddPlace(*testPlace)
	assert.NoError(err)
}

func TestTheaterData_AddPlaceErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPlace)).
		WithArgs(testPlace.SectorId, testPlace.Name).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddPlace(*testPlace)
	assert.Error(err)
}

func TestTheaterData_AddPoster(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPoster)).
		WithArgs(testPoster1.AccountId, testPoster1.ScheduleId, testPoster1.Comment).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddPoster(*testPoster1)
	assert.NoError(err)
}

func TestTheaterData_AddPosterErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPoster)).
		WithArgs(testPoster1.AccountId, testPoster1.ScheduleId, testPoster1.Comment).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddPoster(*testPoster1)
	assert.Error(err)
}

func TestTheaterData_AddPrice(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPrice)).
		WithArgs(testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddPrice(*testPrice)
	assert.NoError(err)
}
func TestTheaterData_AddPriceErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertPrice)).
		WithArgs(testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddPrice(*testPrice)
	assert.Error(err)
}

func TestTheaterData_AddRole(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertRole)).
		WithArgs(testRole.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddRole(*testRole)
	assert.NoError(err)
}

func TestTheaterData_AddRoleErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertRole)).
		WithArgs(testRole.Name).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddRole(*testRole)
	assert.Error(err)
}

func TestTheaterData_AddSchedule(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertSchedule)).
		WithArgs(testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddSchedule(*testSchedule)
	assert.NoError(err)
}

func TestTheaterData_AddScheduleErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertSchedule)).
		WithArgs(testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddSchedule(*testSchedule)
	assert.Error(err)
}

func TestTheaterData_AddSector(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertSector)).
		WithArgs(testSector.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddSector(*testSector)
	assert.NoError(err)
}

func TestTheaterData_AddSectorErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertSector)).
		WithArgs(testSector.Name).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddSector(*testSector)
	assert.Error(err)
}

func TestTheaterData_AddTicket(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertTicket)).
		WithArgs(testTicket1.AccountId, testTicket1.ScheduleId,
			testTicket1.PlaceId, testTicket1.DateOfIssue, testTicket1.Paid,
			testTicket1.Reservation, testTicket1.Destroyed).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddTicket(*testTicket1)
	assert.NoError(err)
}

func TestTheaterData_AddTicketErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertTicket)).
		WithArgs(testTicket1.AccountId, testTicket1.ScheduleId,
			testTicket1.PlaceId, testTicket1.DateOfIssue, testTicket1.Paid,
			testTicket1.Reservation, testTicket1.Destroyed).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddTicket(*testTicket1)
	assert.Error(err)
}

func TestTheaterData_AddUser(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertUser)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.AddUser(*testUser1)
	assert.NoError(err)
}

func TestTheaterData_AddUserErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(insertUser)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber).
		WillReturnError(errors.New("something went wrong..."))
	err := data.AddUser(*testUser1)
	assert.Error(err)
}

func TestTheaterData_DeleteEntry(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(deleteBegin + testTable.String() + deleteEnd)).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.DeleteEntry(testTable, id)
	assert.NoError(err)
}

func TestTheaterData_DeleteEntryErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(deleteBegin + testTable.String() + deleteEnd)).
		WithArgs(id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.DeleteEntry(testTable, id)
	assert.Error(err)
}

func TestTheaterData_UpdateAccount(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateAccount)).
		WithArgs(testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email, testAccount.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateAccount(*testAccount)
	assert.NoError(err)
}

func TestTheaterData_UpdateAccountErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateAccount)).
		WithArgs(testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email, testAccount.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateAccount(*testAccount)
	assert.Error(err)
}

func TestTheaterData_UpdateGenre(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateGenre)).
		WithArgs(testGenre.Name, testGenre.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateGenre(*testGenre)
	assert.NoError(err)
}

func TestTheaterData_UpdateGenreErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateGenre)).
		WithArgs(testGenre.Name, testGenre.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateGenre(*testGenre)
	assert.Error(err)
}

func TestTheaterData_UpdateHall(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateHall)).
		WithArgs(testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId, testHall.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateHall(*testHall)
	assert.NoError(err)
}

func TestTheaterData_UpdateHallErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateHall)).
		WithArgs(testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId, testHall.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateHall(*testHall)
	assert.Error(err)
}

func TestTheaterData_UpdateLocation(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateLocation)).
		WithArgs(testLocation.AccountId, testLocation.Address,
			testLocation.PhoneNumber, testLocation.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateLocation(*testLocation)
	assert.NoError(err)
}

func TestTheaterData_UpdateLocationErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateLocation)).
		WithArgs(testLocation.AccountId, testLocation.Address,
			testLocation.PhoneNumber, testLocation.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateLocation(*testLocation)
	assert.Error(err)
}

func TestTheaterData_UpdatePerformance(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePerformance)).
		WithArgs(testPerformance.AccountId, testPerformance.Name,
			testPerformance.GenreId, testPerformance.Duration, testPerformance.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdatePerformance(*testPerformance)
	assert.NoError(err)
}

func TestTheaterData_UpdatePerformanceErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePerformance)).
		WithArgs(testPerformance.AccountId, testPerformance.Name,
			testPerformance.GenreId, testPerformance.Duration, testPerformance.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdatePerformance(*testPerformance)
	assert.Error(err)
}

func TestTheaterData_UpdatePlace(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePlace)).
		WithArgs(testPlace.SectorId, testPlace.Name, testPlace.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdatePlace(*testPlace)
	assert.NoError(err)
}

func TestTheaterData_UpdatePlaceErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePlace)).
		WithArgs(testPlace.SectorId, testPlace.Name, testPlace.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdatePlace(*testPlace)
	assert.Error(err)
}

func TestTheaterData_UpdatePoster(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePoster)).
		WithArgs(testPoster1.AccountId, testPoster1.ScheduleId,
			testPoster1.Comment, testPoster.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdatePoster(*testPoster1)
	assert.NoError(err)
}

func TestTheaterData_UpdatePosterErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePoster)).
		WithArgs(testPoster1.AccountId, testPoster1.ScheduleId,
			testPoster1.Comment, testPoster.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdatePoster(*testPoster1)
	assert.Error(err)
}

func TestTheaterData_UpdatePrice(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePrice)).
		WithArgs(testPrice.AccountId, testPrice.SectorId,
			testPrice.PerformanceId, testPrice.Price, testPrice.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdatePrice(*testPrice)
	assert.NoError(err)
}
func TestTheaterData_UpdatePriceErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updatePrice)).
		WithArgs(testPrice.AccountId, testPrice.SectorId,
			testPrice.PerformanceId, testPrice.Price, testPrice.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdatePrice(*testPrice)
	assert.Error(err)
}

func TestTheaterData_UpdateRole(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateRole)).
		WithArgs(testRole.Name, testRole.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateRole(*testRole)
	assert.NoError(err)
}

func TestTheaterData_UpdateRoleErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateRole)).
		WithArgs(testRole.Name, testRole.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateRole(*testRole)
	assert.Error(err)
}

func TestTheaterData_UpdateSchedule(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateSchedule)).
		WithArgs(testSchedule.AccountId, testSchedule.PerformanceId,
			testSchedule.Date, testSchedule.HallId, testSchedule.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateSchedule(*testSchedule)
	assert.NoError(err)
}

func TestTheaterData_UpdateScheduleErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateSchedule)).
		WithArgs(testSchedule.AccountId, testSchedule.PerformanceId,
			testSchedule.Date, testSchedule.HallId, testSchedule.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateSchedule(*testSchedule)
	assert.Error(err)
}

func TestTheaterData_UpdateSector(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateSector)).
		WithArgs(testSector.Name, testSector.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateSector(*testSector)
	assert.NoError(err)
}

func TestTheaterData_UpdateSectorErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateSector)).
		WithArgs(testSector.Name, testSector.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateSector(*testSector)
	assert.Error(err)
}

func TestTheaterData_UpdateTicket(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateTicket)).
		WithArgs(testTicket1.AccountId, testTicket1.ScheduleId,
			testTicket1.PlaceId, testTicket1.DateOfIssue, testTicket1.Paid,
			testTicket1.Reservation, testTicket1.Destroyed, testTicket1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateTicket(*testTicket1)
	assert.NoError(err)
}

func TestTheaterData_UpdateTicketErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateTicket)).
		WithArgs(testTicket1.AccountId, testTicket1.ScheduleId,
			testTicket1.PlaceId, testTicket1.DateOfIssue, testTicket1.Paid,
			testTicket1.Reservation, testTicket1.Destroyed, testTicket1.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateTicket(*testTicket1)
	assert.Error(err)
}

func TestTheaterData_UpdateUser(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateUser)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber, testUser1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.UpdateUser(*testUser1)
	assert.NoError(err)
}

func TestTheaterData_UpdateUserErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateUser)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber, testUser1.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.UpdateUser(*testUser1)
	assert.Error(err)
}
