package tests

import (
	"context"
	"database/sql"
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/tmammado/take-home-assignment/model"
	"github.com/tmammado/take-home-assignment/repository"
)

var _ = Describe("User", func() {
	var userRepo *repository.UserRepo
	var mock sqlmock.Sqlmock
	email := "tahir@gmail.com"
	user := &model.User{
		Email:     email,
		Password:  "password",
		FirstName: "Tahir",
		LastName:  "Mammadov",
	}
	user2 := &model.User{
		Email:     "tahir1@gmail.com",
		Password:  "password",
		FirstName: "Tahir",
		LastName:  "Mammadov",
	}

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New() 
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}) 
		Expect(err).ShouldNot(HaveOccurred())

		userRepo = repository.NewUserRepo(gdb)
	})
	AfterEach(func() {
		err := mock.ExpectationsWereMet() 
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("GetUserByEmail", func() {
		It("should return user", func() {
			rows := sqlmock.
				NewRows([]string{"email", "password", "first_name", "last_name"}).
				AddRow(user.Email, user.Password, user.FirstName, user.LastName)

			const sqlSelectOne = `SELECT * FROM "users" WHERE "users"."email" = $1 ORDER BY "users"."email" LIMIT 1`

			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
				WithArgs(email).
				WillReturnRows(rows)

			dbUser, err := userRepo.GetUserByEmail(context.TODO(), email)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(*dbUser).Should(Equal(*user))
		})

		It("should return error", func() {
			const sqlSelectOne = `SELECT * FROM "users" WHERE "users"."email" = $1 ORDER BY "users"."email" LIMIT 1`

			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
				WithArgs(email).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := userRepo.GetUserByEmail(context.TODO(), email)
			Expect(err).Should(HaveOccurred())
		})

	})
	Context("GetAllUsers", func() {
		It("should all users in DB", func() {
			rows := sqlmock.
				NewRows([]string{"email", "password", "first_name", "last_name"}).
				AddRow(user.Email, user.Password, user.FirstName, user.LastName).
				AddRow(user2.Email, user2.Password, user2.FirstName, user2.LastName)

			const sqlSelect = `SELECT * FROM "users"`

			mock.ExpectQuery(regexp.QuoteMeta(sqlSelect)).
				WillReturnRows(rows)

			dbUsers, err := userRepo.GetAllUsers(context.TODO())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(*dbUsers)).Should(Equal(2))
		})
	})
})
