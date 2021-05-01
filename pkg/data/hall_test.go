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

type HallSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *HallData
}

func TestInitHall(t *testing.T) {
	suite.Run(t, new(HallSuite))
}

func (s *HallSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *HallSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewHallData(s.DB)
}

var testHall = &Hall{
	Id:         4,
	AccountId:  1,
	Name:       "testName",
	Capacity:   1099,
	LocationId: 1,
}

func (s *HallSuite) TestTheaterData_AddHall() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "halls"`)).
		WithArgs(testHall.Id, testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddHall(*testHall)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *HallSuite) TestTheaterData_AddHallErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "halls"`)).
		WithArgs(testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddHall(*testHall)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *HallSuite) TestTheaterData_DeleteHall() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "halls"`)).
		WithArgs(testHall.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteHall(*testHall)
	require.NoError(s.T(), err)
}

func (s *HallSuite) TestTheaterData_DeleteHallErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "halls"`)).
		WithArgs(testHall.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteHall(*testHall)
	require.Error(s.T(), err)
}

func (s *HallSuite) TestTheaterData_UpdateHall() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "halls"`)).
		WithArgs(testHall.AccountId, testHall.Capacity, testHall.Id,
			testHall.LocationId, testHall.Name, testHall.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateHall(*testHall)
	require.NoError(s.T(), err)
}

func (s *HallSuite) TestTheaterData_UpdateHallErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "halls"`)).
		WithArgs(testHall.AccountId, testHall.Capacity, testHall.Id,
			testHall.LocationId, testHall.Name, testHall.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateHall(*testHall)
	require.Error(s.T(), err)
}

func (s *HallSuite) TestTheaterData_FindByIdHall() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "name", "capacity", "location_id"}).
		AddRow(testHall.Id, testHall.AccountId, testHall.Name, testHall.Capacity, testHall.LocationId)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "halls" WHERE "halls"."id" = $1 ORDER BY "halls"."id" ASC LIMIT 1`)).
		WithArgs(testHall.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdHall(Hall{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testHall, &res))
}

func (s *HallSuite) TestTheaterData_FindByIdHallErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "halls" WHERE "halls"."id" = $1 ORDER BY "halls"."id" ASC LIMIT 1`)).
		WithArgs(testHall.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdHall(Hall{Id: 4})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
