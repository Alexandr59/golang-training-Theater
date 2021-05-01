package data

import (
	"database/sql"
	"errors"
	"github.com/go-test/deep"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SectorSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *SectorData
}

func TestInitSector(t *testing.T) {
	suite.Run(t, new(SectorSuite))
}

func (s *SectorSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *SectorSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewSectorData(s.DB)
}

var testSector = &Sector{
	Id:   9,
	Name: "L",
}

func (s *SectorSuite) TestTheaterData_AddSector() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "sectors"`)).
		WithArgs(testSector.Id, testSector.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddSector(*testSector)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *SectorSuite) TestTheaterData_AddSectorErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "sectors"`)).
		WithArgs(testSector.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddSector(*testSector)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *SectorSuite) TestTheaterData_DeleteSector() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "sectors"`)).
		WithArgs(testSector.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteSector(*testSector)
	require.NoError(s.T(), err)
}

func (s *SectorSuite) TestTheaterData_DeleteSectorErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "sectors"`)).
		WithArgs(testSector.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteSector(*testSector)
	require.Error(s.T(), err)
}

func (s *SectorSuite) TestTheaterData_UpdateSector() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "sectors"`)).
		WithArgs(testSector.Id, testSector.Name, testSector.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateSector(*testSector)
	require.NoError(s.T(), err)
}

func (s *SectorSuite) TestTheaterData_UpdateSectorErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "sectors"`)).
		WithArgs(testSector.Id, testSector.Name, testSector.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateSector(*testSector)
	require.Error(s.T(), err)
}

func (s *SectorSuite) TestTheaterData_FindByIdSector() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(testSector.Id, testSector.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "sectors" WHERE "sectors"."id" = $1 ORDER BY "sectors"."id" ASC LIMIT 1`)).
		WithArgs(testSector.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdSector(Sector{Id: 9})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testSector, &res))
}

func (s *SectorSuite) TestTheaterData_FindByIdSectorErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "sectors" WHERE "sectors"."id" = $1 ORDER BY "sectors"."id" ASC LIMIT 1`)).
		WithArgs(testSector.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdSector(Sector{Id: 9})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
