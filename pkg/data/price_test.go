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

type PriceSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *PriceData
}

func TestInitPrice(t *testing.T) {
	suite.Run(t, new(PriceSuite))
}

func (s *PriceSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *PriceSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewPriceData(s.DB)
}

var testPrice = &Price{
	Id:            6,
	AccountId:     1,
	SectorId:      10,
	PerformanceId: 2,
	Price:         120,
}

func (s *PriceSuite) TestTheaterData_AddPrice() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "prices"`)).
		WithArgs(testPrice.Id, testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddPrice(*testPrice)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}
func (s *PriceSuite) TestTheaterData_AddPriceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "prices"`)).
		WithArgs(testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddPrice(*testPrice)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *PriceSuite) TestTheaterData_DeletePrice() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "prices"`)).
		WithArgs(testPrice.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeletePrice(*testPrice)
	require.NoError(s.T(), err)
}

func (s *PriceSuite) TestTheaterData_DeletePriceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "prices"`)).
		WithArgs(testPrice.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeletePrice(*testPrice)
	require.Error(s.T(), err)
}

func (s *PriceSuite) TestTheaterData_UpdatePrice() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "prices"`)).
		WithArgs(testPrice.AccountId, testPrice.Id, testPrice.PerformanceId,
			testPrice.Price, testPrice.SectorId, testPrice.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdatePrice(*testPrice)
	require.NoError(s.T(), err)
}

func (s *PriceSuite) TestTheaterData_UpdatePriceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "prices"`)).
		WithArgs(testPrice.AccountId, testPrice.Id, testPrice.PerformanceId,
			testPrice.Price, testPrice.SectorId, testPrice.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdatePrice(*testPrice)
	require.Error(s.T(), err)
}

func (s *PriceSuite) TestTheaterData_FindByIdPrice() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "sector_id", "performance_id", "price"}).
		AddRow(testPrice.Id, testPrice.AccountId, testPrice.SectorId, testPrice.PerformanceId, testPrice.Price)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "prices" WHERE "prices"."id" = $1 ORDER BY "prices"."id" ASC LIMIT 1`)).
		WithArgs(testPrice.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdPrice(Price{Id: 6})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPrice, &res))
}

func (s *PriceSuite) TestTheaterData_FindByIdPriceErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "prices" WHERE "prices"."id" = $1 ORDER BY "prices"."id" ASC LIMIT 1`)).
		WithArgs(testPrice.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdPrice(Price{Id: 6})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
