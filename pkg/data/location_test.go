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

type LocationSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *LocationData
}

func TestInitLocation(t *testing.T) {
	suite.Run(t, new(LocationSuite))
}

func (s *LocationSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *LocationSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewLocationData(s.DB)
}

var testLocation = &Location{
	Id:          4,
	AccountId:   1,
	Address:     "Gaidara10",
	PhoneNumber: "+3754466633321",
}

func (s *LocationSuite) TestTheaterData_AddLocation() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "locations"`)).
		WithArgs(testLocation.Id, testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddLocation(*testLocation)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *LocationSuite) TestTheaterData_AddLocationErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "locations"`)).
		WithArgs(testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddLocation(*testLocation)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *LocationSuite) TestTheaterData_DeleteLocation() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "locations"`)).
		WithArgs(testLocation.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteLocation(*testLocation)
	require.NoError(s.T(), err)
}

func (s *LocationSuite) TestTheaterData_DeleteLocationErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "locations"`)).
		WithArgs(testLocation.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteLocation(*testLocation)
	require.Error(s.T(), err)
}

func (s *LocationSuite) TestTheaterData_UpdateLocation() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "locations"`)).
		WithArgs(testLocation.AccountId, testLocation.Address,
			testLocation.Id, testLocation.PhoneNumber, testLocation.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateLocation(*testLocation)
	require.NoError(s.T(), err)
}

func (s *LocationSuite) TestTheaterData_UpdateLocationErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "locations"`)).
		WithArgs(testLocation.AccountId, testLocation.Address,
			testLocation.Id, testLocation.PhoneNumber, testLocation.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateLocation(*testLocation)
	require.Error(s.T(), err)
}

func (s *LocationSuite) TestTheaterData_FindByIdLocation() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "address", "phone_number"}).
		AddRow(testLocation.Id, testLocation.AccountId, testLocation.Address, testLocation.PhoneNumber)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "locations" WHERE "locations"."id" = $1 ORDER BY "locations"."id" ASC LIMIT 1`)).
		WithArgs(testLocation.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdLocation(Location{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testLocation, &res))
}

func (s *LocationSuite) TestTheaterData_FindByIdLocationErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "locations" WHERE "locations"."id" = $1 ORDER BY "locations"."id" ASC LIMIT 1`)).
		WithArgs(testLocation.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdLocation(Location{Id: 4})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
