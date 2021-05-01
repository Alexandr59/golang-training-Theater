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

type UserSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *UserData
}

func TestInitUser(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *UserSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewUserData(s.DB)
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

var testUser = &SelectUser{
	Id:                  1,
	FirstName:           "Charles",
	LastName:            "Dean",
	Role:                "Actor",
	LocationAddress:     "Gaidara_6",
	LocationPhoneNumber: "+375443564987",
	PhoneNumber:         "+375445239375",
}

func (s *UserSuite) TestTheaterData_AddUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs(testUser1.Id, testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddUser(*testUser1)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *UserSuite) TestTheaterData_AddUserErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.LastName,
			testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddUser(*testUser1)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *UserSuite) TestTheaterData_DeleteUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"`)).
		WithArgs(testUser1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteUser(*testUser1)
	require.NoError(s.T(), err)
}

func (s *UserSuite) TestTheaterData_DeleteUserErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"`)).
		WithArgs(testUser1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteUser(*testUser1)
	require.Error(s.T(), err)
}

func (s *UserSuite) TestTheaterData_UpdateUser() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.Id, testUser1.LastName,
			testUser1.LocationId, testUser1.PhoneNumber, testUser1.RoleId, testUser1.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateUser(*testUser1)
	require.NoError(s.T(), err)
}

func (s *UserSuite) TestTheaterData_UpdateUserErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(testUser1.AccountId, testUser1.FirstName, testUser1.Id, testUser1.LastName,
			testUser1.LocationId, testUser1.PhoneNumber, testUser1.RoleId, testUser1.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateUser(*testUser1)
	require.Error(s.T(), err)
}

func (s *UserSuite) TestTheaterData_FindByIdUser() {
	rows := sqlmock.NewRows([]string{"id", "account_id", "first_name", "last_name",
		"role_id", "location_id", "phone_number"}).
		AddRow(testUser1.Id, testUser1.AccountId, testUser1.FirstName,
			testUser1.LastName, testUser1.RoleId, testUser1.LocationId, testUser1.PhoneNumber)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(testUser1.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdUser(User{Id: 3})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testUser1, &res))
}

func (s *UserSuite) TestTheaterData_FindByIdUserErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(testUser1.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdUser(User{Id: 3})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *UserSuite) TestTheaterData_ReadAllUsers() {
	rows := sqlmock.NewRows([]string{"users.id", "users.first_name", "users.last_name", "roles.name", "locations.address",
		"locations.phone_number", "users.phone_number"}).
		AddRow(testUser.Id, testUser.FirstName, testUser.LastName, testUser.Role,
			testUser.LocationAddress, testUser.LocationPhoneNumber, testUser.PhoneNumber)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.first_name, users.last_name, roles.name, locations.address, 
locations.phone_number, users.phone_number FROM "users" 
JOIN roles on users.role_id = roles.id 
JOIN locations on locations.id = users.account_id 
WHERE (users.account_id = $1)`)).
		WithArgs(Account{Id: 1}.Id).
		WillReturnRows(rows)
	res, err := s.data.ReadAllUsers(Account{Id: 1})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testUser, &res[0]))
}

func (s *UserSuite) TestTheaterData_ReadAllUsersErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.first_name, users.last_name, roles.name, locations.address, 
locations.phone_number, users.phone_number FROM "users" 
JOIN roles on users.role_id = roles.id 
JOIN locations on locations.id = users.account_id 
WHERE (users.account_id = $1)`)).
		WithArgs(Account{Id: 1}.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.ReadAllUsers(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
