package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tmammado/take-home-assignment/model"
)

var _ = Describe("App", func() {
	var user model.User
	var password = "12345678"

	BeforeEach(func() {
		user = model.User{
			Email:     "tahir@gmail.com",
			Password:  password,
			FirstName: "Tahir",
			LastName:  "Mammadov",
		}
	})

	Context("HashPassword", func() {
		It("should hash user password", func() {
			user.HashPassword()

			Expect(user.Password).ShouldNot(Equal(password))
		})
	})

	Context("CheckPassword", func() {
		It("should return error when password is not correct", func() {
			user.HashPassword()
			err := user.CheckPassword("123456")

			Expect(err).Should(HaveOccurred())
		})

		It("should return nil when password is correct", func() {
			user.HashPassword()
			err := user.CheckPassword(password)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
