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

type PlaceSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *PlaceData
}

func TestInitPlace(t *testing.T) {
	suite.Run(t, new(PlaceSuite))
}

func (s *PlaceSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *PlaceSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewPlaceData(s.DB)
}

var testPlace = &Place{
	Id:       6,
	SectorId: 15,
	Name:     "2",
}

func (s *PlaceSuite) TestTheaterData_AddPlace() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "places"`)).
		WithArgs(testPlace.Id, testPlace.SectorId, testPlace.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPlace(*testPlace)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *PlaceSuite) TestTheaterData_AddPlaceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "places"`)).
		WithArgs(testPlace.SectorId, testPlace.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPlace(*testPlace)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *PlaceSuite) TestTheaterData_DeletePlace() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "places"`)).
		WithArgs(testPlace.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePlace(*testPlace)
	require.NoError(s.T(), err)
}

func (s *PlaceSuite) TestTheaterData_DeletePlaceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "places"`)).
		WithArgs(testPlace.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePlace(*testPlace)
	require.Error(s.T(), err)
}

func (s *PlaceSuite) TestTheaterData_UpdatePlace() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "places"`)).
		WithArgs(testPlace.Id, testPlace.Name, testPlace.SectorId, testPlace.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePlace(*testPlace)
	require.NoError(s.T(), err)
}

func (s *PlaceSuite) TestTheaterData_UpdatePlaceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "places"`)).
		WithArgs(testPlace.Id, testPlace.Name, testPlace.SectorId, testPlace.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePlace(*testPlace)
	require.Error(s.T(), err)
}

func (s *PlaceSuite) TestTheaterData_FindByIdPlace() {
	rows := sqlmock.NewRows([]string{"id", "sector_id", "name"}).
		AddRow(testPlace.Id, testPlace.SectorId, testPlace.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "places" WHERE "places"."id" = $1 ORDER BY "places"."id" ASC LIMIT 1`)).
		WithArgs(testPlace.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPlace(Place{Id: 6})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPlace, &res))
}

func (s *PlaceSuite) TestTheaterData_FindByIdPlaceErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "places" WHERE "places"."id" = $1 ORDER BY "places"."id" ASC LIMIT 1`)).
		WithArgs(testPlace.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPlace(Place{Id: 6})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
