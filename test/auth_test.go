package test

import (
	"net/http"
	"testing"

	"github.com/akifkadioglu/youtube-1/pkg/auth"
	"github.com/akifkadioglu/youtube-1/route"
)

func TestAuth(t *testing.T) {
	setupTest()
	s := route.CreateServer()

	t.Run("Register", func(t *testing.T) {

		body := auth.BodyRegister{
			Username: "akif",
			Password: "12345",
		}
		req, _ := http.NewRequest("POST", "/register", setBodyForTest(body))
		response := executeRequest(req, s)
		checkResponseCode(t, http.StatusOK, response.Code)
	})

	t.Run("Login", func(t *testing.T) {

		body := auth.BodyLogin{
			Username: "akif",
			Password: "12345",
		}
		req, _ := http.NewRequest("POST", "/login", setBodyForTest(body))
		response := executeRequest(req, s)
		checkResponseCode(t, http.StatusOK, response.Code)
	})
}
