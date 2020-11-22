package monit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStatusOK(t *testing.T) {
	ut := UserHTTPTest{
		URL:      "http://localhost:8080",
		Status:   200,
		Interval: time.Minute,
	}

	expectedStatus := serviceStatus{status: "Down"}
	serviceStatus := RunUserTest(ut)
	require.Equal(t, expectedStatus, serviceStatus)
}
