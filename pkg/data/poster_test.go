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

type PosterSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *PosterData
}

func TestInitPoster(t *testing.T) {
	suite.Run(t, new(PosterSuite))
}

func (s *PosterSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *PosterSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewPosterData(s.DB)
}

var testPoster = &Poster{
	Id:         2,
	AccountId:  1,
	ScheduleId: 1,
	Comment:    "Hi!!!",
}

var testSelectPoster = &SelectPoster{
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

func (s *PosterSuite) TestTheaterData_AddPoster() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "posters"`)).
		WithArgs(testPoster.Id, testPoster.AccountId, testPoster.ScheduleId, testPoster.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPoster(*testPoster)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *PosterSuite) TestTheaterData_AddPosterErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "posters"`)).
		WithArgs(testPoster.AccountId, testPoster.ScheduleId, testPoster.Comment).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPoster(*testPoster)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *PosterSuite) TestTheaterData_DeletePoster() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "posters"`)).
		WithArgs(testPoster.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePoster(*testPoster)
	require.NoError(s.T(), err)
}

func (s *PosterSuite) TestTheaterData_DeletePosterErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "posters"`)).
		WithArgs(testPoster.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePoster(*testPoster)
	require.Error(s.T(), err)
}

func (s *PosterSuite) TestTheaterData_UpdatePoster() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posters"`)).
		WithArgs(testPoster.AccountId, testPoster.Comment,
			testPoster.Id, testPoster.ScheduleId, testPoster.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePoster(*testPoster)
	require.NoError(s.T(), err)
}

func (s *PosterSuite) TestTheaterData_UpdatePosterErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posters"`)).
		WithArgs(testPoster.AccountId, testPoster.Comment,
			testPoster.Id, testPoster.ScheduleId, testPoster.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePoster(*testPoster)
	require.Error(s.T(), err)
}

func (s *PosterSuite) TestTheaterData_FindByIdPoster() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "schedule_id", "comment"}).
		AddRow(testPoster.Id, testPoster.AccountId, testPoster.ScheduleId, testPoster.Comment)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posters" WHERE "posters"."id" = $1 ORDER BY "posters"."id" ASC LIMIT 1`)).
		WithArgs(testPoster.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPoster(Poster{Id: 2})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPoster, &res))
}

func (s *PosterSuite) TestTheaterData_FindByIdPosterErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posters" WHERE "posters"."id" = $1 ORDER BY "posters"."id" ASC LIMIT 1`)).
		WithArgs(testPoster.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPoster(Poster{Id: 2})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *PosterSuite) TestTheaterData_ReadAllPosters() {
	rows := sqlmock.NewRows([]string{"poster.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "poster.comment"}).
		AddRow(testSelectPoster.Id, testSelectPoster.PerformanceName, testSelectPoster.GenreName, testSelectPoster.PerformanceDuration,
			testSelectPoster.DateTime, testSelectPoster.HallName, testSelectPoster.HallCapacity, testSelectPoster.LocationAddress,
			testSelectPoster.LocationPhoneNumber, testSelectPoster.Comment)
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
	require.Nil(s.T(), deep.Equal(testSelectPoster, &res[0]))
}

func (s *PosterSuite) TestTheaterData_ReadAllPostersErr() {
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
