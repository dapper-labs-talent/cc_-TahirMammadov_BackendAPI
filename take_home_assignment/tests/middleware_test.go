package tests

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tmammado/take-home-assignment/middleware"
)

var _ = Describe("App", func() {
	var res *httptest.ResponseRecorder
	var validJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODcyMTE0OTAsIkVtYWlsIjoidGFoaXJAZ21haWwuY29tIn0._tbeEOdc308B3Hq25WKM6kV_bR0UiLcpB0Iw4cTKfCI"
	var invalidJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTU2NzUwNTYsIkVtYWlsIjoidGFoaXJAZ21haWwuY29tIn0.j1A2Mn09PmFISAfhMNCn53VuonmMAmDmY0JrQ8KBqWU"
	var handler = func(w http.ResponseWriter, r *http.Request) {
		return
	}

	BeforeEach(func() {
		res = httptest.NewRecorder()
	})

	Context("JwtAuthentication", func() {
		It("should return 401 response with header has no JWT", func() {
			req, _ := http.NewRequest("GET", "/users", nil)
			middleware.JwtAuthentication(http.HandlerFunc(handler)).ServeHTTP(res, req)

			Expect(res.Code).Should(Equal(http.StatusUnauthorized))
		})

		It("should return 401 response if JWT is invalid", func() {
			req, _ := http.NewRequest("GET", "/users", nil)
			req.Header.Set("x-authentication-token", invalidJWT)
			middleware.JwtAuthentication(http.HandlerFunc(handler)).ServeHTTP(res, req)

			Expect(res.Code).Should(Equal(http.StatusUnauthorized))
		})

		It("should return 200 response if JWT is valid", func() {
			req, _ := http.NewRequest("GET", "/users", nil)
			req.Header.Set("x-authentication-token", validJWT)
			middleware.JwtAuthentication(http.HandlerFunc(handler)).ServeHTTP(res, req)

			Expect(res.Code).Should(Equal(http.StatusOK))
		})
	})
})
