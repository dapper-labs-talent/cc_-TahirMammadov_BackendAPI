package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tmammado/take-home-assignment/handler"
	"github.com/tmammado/take-home-assignment/model"
	"github.com/tmammado/take-home-assignment/shared"
)

var _ = Describe("App", func() {
	var app *handler.App
	var appDefault *handler.App
	var res *httptest.ResponseRecorder

	BeforeEach(func() {
		user := model.User{Email: "tahir@gmail.com", FirstName: "Tahir", LastName: "Mammadov", Password: "12345678"}
		mockUserRepo := MockUserRepo{user: user}
		mockUserRepoDefault := MockUserRepoDefault{user: user}
		app = handler.NewApp(&mockUserRepo)
		appDefault = handler.NewApp(&mockUserRepoDefault)
		res = httptest.NewRecorder()
	})

	Context("Signup", func() {
		It("should return 200 response", func() {
			var jsonStr = []byte(`{
				"email": "tahir@gmail.com",
				"password": "12345678",
				"firstName": "Tahir",
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))

			app.Signup(res, req)

			var resBody shared.TokenResponse
			if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
				log.Fatalln(err)
			}

			Expect(res.Code).Should(Equal(http.StatusOK))
			Expect(resBody.Token).ShouldNot(BeNil())
		})

		It("should return 400 response if user exists", func() {
			var jsonStr = []byte(`{
				"email": "tahir@gmail.com",
				"password": "12345678",
				"firstName": "Tahir",
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))

			appDefault.Signup(res, req)

			Expect(res.Code).Should(Equal(http.StatusBadRequest))
		})

		It("should return 400 response if request validation fails", func() {
			var jsonStr = []byte(`{
				"email": "tahir@gmail.com",
				"password": "12345",
				"firstName": "Tahir",
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))

			app.Signup(res, req)

			Expect(res.Code).Should(Equal(http.StatusBadRequest))
		})

		It("should return 400 response if it cannot decode req body", func() {
			var jsonStr = []byte(`{
				"email": "tahir@gmail.com",
				"password": 1,
				"firstName": "Tahir",
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))

			app.Signup(res, req)

			Expect(res.Code).Should(Equal(http.StatusBadRequest))
		})
	})

	Context("Login", func() {
		It("should return 200 response", func() {
			var jsonStr = []byte(`{
				"email": "tahir@gmail.com",
				"password": "12345678"
			  }`)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))

			appDefault.Login(res, req)

			var resBody shared.TokenResponse
			if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
				log.Fatalln(err)
			}

			Expect(res.Code).Should(Equal(http.StatusOK))
			Expect(resBody.Token).ShouldNot(BeNil())
		})

		It("should return 400 response if request validation fails", func() {
			var jsonStr = []byte(`{
				"email": "tahir@gmail.com",
				"password": "12345",
			  }`)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))

			app.Login(res, req)

			Expect(res.Code).Should(Equal(http.StatusBadRequest))
		})

		It("should return 400 response if it cannot decode req body", func() {
			var jsonStr = []byte(`{
				"email": "tahir@gmail.com",
				"password": 1,
			  }`)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))

			app.Login(res, req)

			Expect(res.Code).Should(Equal(http.StatusBadRequest))
		})

	})

	Context("GetAllUsers", func() {
		It("should return 200 response", func() {
			var jsonStr = []byte(`{
				"firstName": "Tahir",
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("GET", "/users", bytes.NewBuffer(jsonStr))

			app.GetAllUsers(res, req)

			var resBody shared.UserResponse
			if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
				log.Fatalln(err)
			}

			Expect(res.Code).Should(Equal(http.StatusOK))
			Expect(resBody.Users).ShouldNot(BeNil())
		})
	})

	Context("UpdateUser", func() {
		It("should return 200 response", func() {
			var jsonStr = []byte(`{
				"firstName": "Tahir",
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer(jsonStr))
			ctx := context.WithValue(req.Context(), "email", "email")
			req = req.WithContext(ctx)

			app.UpdateUser(res, req)

			Expect(res.Code).Should(Equal(http.StatusOK))
		})

		It("should return 400 response if request validation fails", func() {
			var jsonStr = []byte(`{
				"firstName": "",
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer(jsonStr))

			app.UpdateUser(res, req)

			Expect(res.Code).Should(Equal(http.StatusBadRequest))
		})

		It("should return 400 response if it cannot decode req body", func() {
			var jsonStr = []byte(`{
				"firstName": 1,
				"lastName": "Mammadov"
			  }`)
			req, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer(jsonStr))

			app.UpdateUser(res, req)

			Expect(res.Code).Should(Equal(http.StatusBadRequest))
		})

	})
})
