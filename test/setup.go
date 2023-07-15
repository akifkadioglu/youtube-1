package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akifkadioglu/youtube-1/database"
	"github.com/akifkadioglu/youtube-1/env"
	"github.com/akifkadioglu/youtube-1/utils"
	"github.com/go-chi/chi/v5"
)

func setupTest() {
	env.InitEnv(env.TEST)
	database.TestConnection()
	utils.InitTokenAuth()
}

func executeRequest(req *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func setBodyForTest(model interface{}) *bytes.Buffer {
	b, _ := json.Marshal(model)
	return bytes.NewBuffer([]byte(b))
}
