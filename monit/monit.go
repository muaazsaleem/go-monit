package monit

import (
	"fmt"
	"time"
)

// userHTTPTest specifies HTTP tests defined by the user
type userHTTPTest struct {
	// URL defines the HTTP endpoint to be tested e.g http://localhost:1235
	URL string `json:"url"`

	// userHTTPTest only supports "GET" for the moment
	// Method string `json:"method"`

	// Status defines the expected HTTP status
	Status string `json:"status"`
	// Interval defines the interval at which the test should be run e.g 20s
	Interval time.Duration `json:"interval"`
}

func Run() {
	fmt.Println("running monit ...")
}
