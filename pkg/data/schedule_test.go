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

type ScheduleSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *ScheduleData
}

func TestInitSchedule(t *testing.T) {
	suite.Run(t, new(ScheduleSuite))
}

func (s *ScheduleSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *ScheduleSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewScheduleData(s.DB)
}

var testSchedule = &Schedule{
	Id:            8,
	AccountId:     1,
	PerformanceId: 3,
	Date:          "2021-04-13 16:00",
	HallId:        3,
}

func (s *ScheduleSuite) TestTheaterData_AddSchedule() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "schedules"`)).
		WithArgs(testSchedule.Id, testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddSchedule(*testSchedule)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *ScheduleSuite) TestTheaterData_AddScheduleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "schedules"`)).
		WithArgs(testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddSchedule(*testSchedule)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *ScheduleSuite) TestTheaterData_DeleteSchedule() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "schedules"`)).
		WithArgs(testSchedule.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteSchedule(*testSchedule)
	require.NoError(s.T(), err)
}

func (s *ScheduleSuite) TestTheaterData_DeleteScheduleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "schedules"`)).
		WithArgs(testSchedule.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteSchedule(*testSchedule)
	require.Error(s.T(), err)
}

func (s *ScheduleSuite) TestTheaterData_UpdateSchedule() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "schedules"`)).
		WithArgs(testSchedule.AccountId, testSchedule.Date, testSchedule.HallId,
			testSchedule.Id, testSchedule.PerformanceId, testSchedule.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateSchedule(*testSchedule)
	require.NoError(s.T(), err)
}

func (s *ScheduleSuite) TestTheaterData_UpdateScheduleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "schedules"`)).
		WithArgs(testSchedule.AccountId, testSchedule.Date, testSchedule.HallId,
			testSchedule.Id, testSchedule.PerformanceId, testSchedule.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateSchedule(*testSchedule)
	require.Error(s.T(), err)
}

func (s *ScheduleSuite) TestTheaterData_FindByIdSchedule() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "performance_id", "date", "hall_id"}).
		AddRow(testSchedule.Id, testSchedule.AccountId, testSchedule.PerformanceId, testSchedule.Date, testSchedule.HallId)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "schedules" WHERE "schedules"."id" = $1 ORDER BY "schedules"."id" ASC LIMIT 1`)).
		WithArgs(testSchedule.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdSchedule(Schedule{Id: 8})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testSchedule, &res))
}

func (s *ScheduleSuite) TestTheaterData_FindByIdScheduleErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "schedules" WHERE "schedules"."id" = $1 ORDER BY "schedules"."id" ASC LIMIT 1`)).
		WithArgs(testSchedule.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdSchedule(Schedule{Id: 8})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
