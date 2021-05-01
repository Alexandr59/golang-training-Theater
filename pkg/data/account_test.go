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

type AccountSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *AccountData
}

func TestInitAccount(t *testing.T) {
	suite.Run(t, new(AccountSuite))
}

func (s *AccountSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *AccountSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewAccountData(s.DB)
}

var testAccount = &Account{
	Id:          4,
	FirstName:   "Dim",
	LastName:    "Ivanov",
	PhoneNumber: "+375296574897",
	Email:       "dimaivanov@gmail.com",
}

func (s *AccountSuite) TestTheaterData_AddAccount() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts"`)).
		WithArgs(testAccount.Id, testAccount.FirstName, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddAccount(*testAccount)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *AccountSuite) TestTheaterData_AddAccountErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts"`)).
		WithArgs(testAccount.Id, testAccount.FirstName, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Email).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddAccount(*testAccount)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *AccountSuite) TestTheaterData_DeleteAccount() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "accounts"`)).
		WithArgs(testAccount.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteAccount(*testAccount)
	require.NoError(s.T(), err)
}

func (s *AccountSuite) TestTheaterData_DeleteAccountErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "accounts"`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteAccount(*testAccount)
	require.Error(s.T(), err)
}

func (s *AccountSuite) TestTheaterData_UpdateAccount() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "accounts"`)).
		WithArgs(testAccount.Email, testAccount.FirstName, testAccount.Id, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateAccount(*testAccount)
	require.NoError(s.T(), err)
}

func (s *AccountSuite) TestTheaterData_UpdateAccountErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "accounts"`)).
		WithArgs(testAccount.Email, testAccount.FirstName, testAccount.Id, testAccount.LastName,
			testAccount.PhoneNumber, testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateAccount(*testAccount)
	require.Error(s.T(), err)
}

func (s *AccountSuite) TestTheaterData_FindByIdAccount() {
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "phone_number", "email"}).
		AddRow(testAccount.Id, testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdAccount(Account{Id: 4})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testAccount, &res))
}

func (s *AccountSuite) TestTheaterData_FindByIdAccountErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdAccount(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
