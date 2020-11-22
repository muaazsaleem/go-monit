package monit

import (
	"time"
)

// UserHTTPTest specifies HTTP tests defined by the user
type UserHTTPTest struct {
	// URL defines the HTTP endpoint to be tested e.g http://localhost:1235
	URL string `json:"url"`

	// UserHTTPTest only supports "GET" for the moment
	// Method string `json:"method"`

	// Status defines the expected HTTP status
	Status int `json:"status"`
	// Interval defines the interval at which the test should be run e.g 20s
	Interval time.Duration `json:"interval"`
}

type serviceStatus struct {
	// status indicates the status of a service i.e "UP", "Failing", "Down"
	status string
}

/* TODO:
 * Take in user test struct
 * generate http request
 * receive response
 * match with test
 * report
 * repeat
 */

func RunUserTest(ut UserHTTPTest) serviceStatus {
	return serviceStatus{status: "UP"}
}
