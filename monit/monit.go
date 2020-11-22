package monit

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// UserHTTPTest specifies HTTP tests defined by the user
type UserHTTPTest struct {
	// URL defines the HTTP endpoint to be tested e.g http://localhost:1235
	URL string `json:"url"`

	// UserHTTPTest only supports "GET" for the moment
	// Method string `json:"method"`

	// StatusCode defines the expected HTTP status
	StatusCode int `json:"status"`
	// Interval defines the interval at which the test should be run e.g 20s
	Interval time.Duration `json:"interval"`
}

// TODO: make this a enum type
// ServiceStatus indicates the status of a service i.e "UP", "FAILING", "DOWN"
type ServiceStatus string

const (
	ServiceUP      = ServiceStatus("UP")
	ServiceDown    = ServiceStatus("DOWN")
	ServiceFailing = ServiceStatus("FAILING")
)

func RunUserTest(ut UserHTTPTest) ServiceStatus {
	// TODO: make sure ut.URL is actually a valid URL
	res, err := http.Get(ut.URL)
	// expecting the HTTP client to err out if the service is down
	// TODO: check or specific errors e.g DNS not resolved
	if err != nil {
		// TODO: add TestURL to the log
		log.Errorf("%s, is 'DOWN'", ut.URL)
		return ServiceDown
	}

	if res.StatusCode == ut.StatusCode {
		log.Infof("%s, is 'UP'", ut.URL)
		return ServiceUP
	} else {
		log.Errorf("%s, is 'FAILING'", ut.URL)
		return ServiceFailing
	}
}
