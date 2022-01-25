package repository

import (
	"bootcampProject/database"
	"bootcampProject/users/domain"
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"testing"
	"time"
)

type SuiteUser struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.UserRepository
	user       *domain.Users
}

func (s *SuiteUser) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(mysql.New(
		mysql.Config{Conn: db,
			SkipInitializeWithVersion: true}),
		&gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 newLogger})
	require.NoError(s.T(), err)

	s.repository = NewUserRepository(database.DBHandler{
		Conn: s.DB})

}

func (s *SuiteUser) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(SuiteUser))
}

func (s *SuiteUser) TestUserRepository_CreateUser() {

	user := domain.Users{
		PwdHash: "test",
		Name:    "nameTest",
		Age:     25,
	}

	idExpected := 1
	/*s.mock.ExpectQuery(
	`INSERT INTO "users" ("name","pwd_hash","age")
		VALUES ($1,$2,$3)`).
	WithArgs(user.Name, user.PwdHash, user.Age).
	WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(idExpected))*/

	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`pwd_hash`,`name`,`age`) VALUES (?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	//WillReturnRows(
	//	sqlmock.NewRows([]string{"id"}).AddRow(idExpected))

	idGot, err := s.repository.CreateUser(context.TODO(), user)
	require.NoError(s.T(), err)
	require.EqualValues(s.T(), idGot, idExpected)
}

func (s *SuiteUser) TestUserRepository_GetUsers() {

	columns := []string{"id", "pwd_hash", "name", "age"}
	row1 := []driver.Value{2, "test1", "nameTest1", 25}
	row2 := []driver.Value{3, "test2", "nameTest2", 30}

	var limit, offset int
	limit = 5
	offset = 1
	s.mock.ExpectQuery("SELECT (.+) FROM `users`").
		//WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(row1...).AddRow(row2...))
	//	WithArgs(limit, offset).
	//WillReturnRows(sqlmock.NewRows([]string{"id, pwd_hash, name, age"}).
	//	AddRow(users[0].ID, users[0].PwdHash, users[0].Name, users[0].Age))
	//	WillReturnRows(
	//		sqlmock.NewRows([]string{"id", "name"}).AddRow(users[0].ID, users[0].Name))

	_, err := s.repository.GetUsers(context.TODO(), limit, offset)
	require.NoError(s.T(), err)
}
