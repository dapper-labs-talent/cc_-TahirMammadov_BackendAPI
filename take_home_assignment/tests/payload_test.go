package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tmammado/take-home-assignment/shared"
)

var _ = Describe("App", func() {
	var createUserRequest shared.CreateUserRequest
	var updateUserRequest shared.UpdateUserRequest
	var credentials shared.Credentials
	var badEmail = "tahirgmail.com"
	var emailWithUppperCase = " Tahir@gmail.com  "
	var badPassword = "1245"
	var badFirstName = ""
	var badLastName = ""
	var stringWithSpaces = "   12345678  "

	BeforeEach(func() {
		createUserRequest = shared.CreateUserRequest{
			Email:     "tahir@gmail.com",
			Password:  "12345678",
			FirstName: "Tahir",
			LastName:  "Mammadov",
		}

		updateUserRequest = shared.UpdateUserRequest{
			FirstName: "Tahir",
			LastName:  "Mammadov",
		}

		credentials = shared.Credentials{
			Email:    "tahir@gmail.com",
			Password: "12345678",
		}
	})

	Context("CreateUserRequest", func() {
		It("should return an error when email is not correct format", func() {
			createUserRequest.Email = badEmail
			err := createUserRequest.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should return an error when password lentgh is less than 8", func() {
			createUserRequest.Password = badPassword
			err := createUserRequest.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should return an error when firstName is empty", func() {
			createUserRequest.FirstName = badFirstName
			err := createUserRequest.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should return an error when lastName is empty", func() {
			createUserRequest.LastName = badLastName
			err := createUserRequest.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should trim spaces", func() {
			createUserRequest.Password = stringWithSpaces
			err := createUserRequest.Validate()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(createUserRequest.Password).ShouldNot(Equal(stringWithSpaces))
		})
	})

	Context("UpdateUserRequest", func() {
		It("should return an error when firstName is empty", func() {
			updateUserRequest.FirstName = badFirstName
			err := updateUserRequest.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should return an error when lastName is empty", func() {
			updateUserRequest.LastName = badLastName
			err := updateUserRequest.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should trim spaces", func() {
			updateUserRequest.FirstName = stringWithSpaces
			updateUserRequest.LastName = stringWithSpaces
			err := updateUserRequest.Validate()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(updateUserRequest.FirstName).ShouldNot(Equal(stringWithSpaces))
			Expect(updateUserRequest.LastName).ShouldNot(Equal(stringWithSpaces))
		})
	})

	Context("Credentials", func() {
		It("should return an error when email is not correct format", func() {
			credentials.Email = badEmail
			err := credentials.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should return an error when password lentgh is less than 8", func() {
			credentials.Password = badPassword
			err := credentials.Validate()

			Expect(err).Should(HaveOccurred())
		})

		It("should trim spaces", func() {
			credentials.Password = stringWithSpaces
			credentials.Email = emailWithUppperCase
			err := credentials.Validate()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(credentials.Password).ShouldNot(Equal(stringWithSpaces))
			Expect(credentials.Email).ShouldNot(Equal(emailWithUppperCase))
		})
	})
})
