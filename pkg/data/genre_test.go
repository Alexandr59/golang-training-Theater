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

type GenreSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *GenreData
}

func TestInitGenre(t *testing.T) {
	suite.Run(t, new(GenreSuite))
}

func (s *GenreSuite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *GenreSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewGenreData(s.DB)
}

var testGenre = &Genre{
	Id:   1,
	Name: "a musical",
}

func (s *GenreSuite) TestTheaterData_AddGenre() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "genres"`)).
		WithArgs(testGenre.Id, testGenre.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
	s.mock.ExpectCommit()
	id, err := s.data.AddGenre(*testGenre)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 4)
}

func (s *GenreSuite) TestTheaterData_AddGenreErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "genres" ("id","name") VALUES ($1,$2) RETURNING "genres"."id"`)).
		WithArgs(testGenre.Name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.data.AddGenre(*testGenre)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *GenreSuite) TestTheaterData_DeleteGenre() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "genres"`)).
		WithArgs(testGenre.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.DeleteGenre(*testGenre)
	require.NoError(s.T(), err)
}

func (s *GenreSuite) TestTheaterData_DeleteGenreErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "genres"`)).
		WithArgs(testGenre.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.DeleteGenre(*testGenre)
	require.Error(s.T(), err)
}

func (s *GenreSuite) TestTheaterData_UpdateGenre() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "genres"`)).
		WithArgs(testGenre.Id, testGenre.Name, testGenre.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.data.UpdateGenre(*testGenre)
	require.NoError(s.T(), err)
}

func (s *GenreSuite) TestTheaterData_UpdateGenreErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "genres"`)).
		WithArgs(testGenre.Id, testGenre.Name, testGenre.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.data.UpdateGenre(*testGenre)
	require.Error(s.T(), err)
}

func (s *GenreSuite) TestTheaterData_FindByIdGenre() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(testGenre.Id, testGenre.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "genres" WHERE "genres"."id" = $1 ORDER BY "genres"."id" ASC LIMIT 1`)).
		WithArgs(testGenre.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdGenre(Genre{Id: 1})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testGenre, &res))
}

func (s *GenreSuite) TestTheaterData_FindByIdGenreErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "genres" WHERE "genres"."id" = $1 ORDER BY "genres"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong"))
	res, err := s.data.FindByIdGenre(Genre{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}
