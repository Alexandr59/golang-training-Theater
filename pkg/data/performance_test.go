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

type PerformanceSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *PerformanceData
}

func TestInitPerformance(t *testing.T) {
	suite.Run(t, new(PerformanceSuite))
}

func (s *PerformanceSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *PerformanceSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewPerformanceData(s.DB)
}

var testPerformance = &Performance{
	Id:        4,
	AccountId: 1,
	Name:      "Big ball",
	GenreId:   3,
	Duration:  "1:00",
}

func (s *PerformanceSuite) TestTheaterData_AddPerformance() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "performances"`)).
		WithArgs(testPerformance.Id, testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPerformance(*testPerformance)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *PerformanceSuite) TestTheaterData_AddPerformanceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "performances"`)).
		WithArgs(testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPerformance(*testPerformance)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *PerformanceSuite) TestTheaterData_DeletePerformance() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "performances"`)).
		WithArgs(testPerformance.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePerformance(*testPerformance)
	require.NoError(s.T(), err)
}

func (s *PerformanceSuite) TestTheaterData_DeletePerformanceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "performances"`)).
		WithArgs(testPerformance.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePerformance(*testPerformance)
	require.Error(s.T(), err)
}

func (s *PerformanceSuite) TestTheaterData_UpdatePerformance() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "performances"`)).
		WithArgs(testPerformance.AccountId, testPerformance.Duration,
			testPerformance.GenreId, testPerformance.Id, testPerformance.Name, testPerformance.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePerformance(*testPerformance)
	require.NoError(s.T(), err)
}

func (s *PerformanceSuite) TestTheaterData_UpdatePerformanceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "performances"`)).
		WithArgs(testPerformance.AccountId, testPerformance.Duration,
			testPerformance.GenreId, testPerformance.Id, testPerformance.Name, testPerformance.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePerformance(*testPerformance)
	require.Error(s.T(), err)
}

func (s *PerformanceSuite) TestTheaterData_FindByIdPerformance() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "name", "genre_id", "duration"}).
		AddRow(testPerformance.Id, testPerformance.AccountId, testPerformance.Name, testPerformance.GenreId, testPerformance.Duration)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "performances" WHERE "performances"."id" = $1 ORDER BY "performances"."id" ASC LIMIT 1`)).
		WithArgs(testPerformance.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPerformance(Performance{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPerformance, &res))
}

func (s *PerformanceSuite) TestTheaterData_FindByIdPerformanceErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "performances" WHERE "performances"."id" = $1 ORDER BY "performances"."id" ASC LIMIT 1`)).
		WithArgs(testPerformance.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPerformance(Performance{Id: 4})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
