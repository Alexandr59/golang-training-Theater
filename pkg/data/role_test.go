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

type RoleSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *RoleData
}

func TestInitRole(t *testing.T) {
	suite.Run(t, new(RoleSuite))
}

func (s *RoleSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *RoleSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewRoleData(s.DB)
}

var testRole = &Role{
	Id:   6,
	Name: "Test",
}

func (s *RoleSuite) TestTheaterData_AddRole() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "roles"`)).
		WithArgs(testRole.Id, testRole.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddRole(*testRole)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *RoleSuite) TestTheaterData_AddRoleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "roles"`)).
		WithArgs(testRole.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddRole(*testRole)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *RoleSuite) TestTheaterData_DeleteRole() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "roles"`)).
		WithArgs(testRole.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteRole(*testRole)
	require.NoError(s.T(), err)
}

func (s *RoleSuite) TestTheaterData_DeleteRoleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "roles"`)).
		WithArgs(testRole.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteRole(*testRole)
	require.Error(s.T(), err)
}

func (s *RoleSuite) TestTheaterData_UpdateRole() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "roles"`)).
		WithArgs(testRole.Id, testRole.Name, testRole.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateRole(*testRole)
	require.NoError(s.T(), err)
}

func (s *RoleSuite) TestTheaterData_UpdateRoleErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "roles"`)).
		WithArgs(testRole.Id, testRole.Name, testRole.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateRole(*testRole)
	require.Error(s.T(), err)
}

func (s *RoleSuite) TestTheaterData_FindByIdRole() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(testRole.Id, testRole.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "roles" WHERE "roles"."id" = $1 ORDER BY "roles"."id" ASC LIMIT 1`)).
		WithArgs(testRole.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdRole(Role{Id: 6})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testRole, &res))
}

func (s *RoleSuite) TestTheaterData_FindByIdRoleErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "roles" WHERE "roles"."id" = $1 ORDER BY "roles"."id" ASC LIMIT 1`)).
		WithArgs(testRole.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdRole(Role{Id: 6})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
