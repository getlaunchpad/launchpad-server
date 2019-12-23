package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/lucasstettner/launchpad-server/app"

	"github.com/joho/godotenv"
)

var a app.App

// This tests the main function of the app, it is necessary that this passes
// Therefore we use the testing.Main package
func TestMain(m *testing.M) {
	if err := godotenv.Load(os.ExpandEnv("../.env")); err != nil {
		log.Printf("Error getting env, continuing in production mode %v\n", err)
	}

	a = app.App{}

	a.Start(false)

	os.Exit(m.Run())
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
