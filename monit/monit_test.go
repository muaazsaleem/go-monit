package monit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var exampleUserTest = UserHTTPTest{
	URL:        "http://localhost:8080",
	StatusCode: 200,
	Interval:   time.Minute,
}

func TestStatusDown(t *testing.T) {
	ss := RunUserTest(exampleUserTest)
	require.Equal(t, ServiceDown, ss)
}

func TestStatusUP(t *testing.T) {
	sv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			_, _ = fmt.Fprint(w, "OK")
		}))
	defer sv.Close()

	// point the test endpoint to the test server
	ut := exampleUserTest
	ut.URL = sv.URL

	ss := RunUserTest(ut)
	require.Equal(t, ServiceUP, ss)
}

func TestStatusFailing(t *testing.T) {
	sv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
	defer sv.Close()

	ut := exampleUserTest
	ut.URL = sv.URL

	ss := RunUserTest(ut)
	require.Equal(t, ServiceFailing, ss)
}
