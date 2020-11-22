package monit

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// HTTPService specifies HTTP tests defined by the user
type HTTPService struct {
	// TODO: include Name here, for more human readable logs
	// URL defines the HTTP endpoint to be tested e.g http://localhost:1235
	URL string `json:"url"`

	// HTTPService only supports "GET" for the moment
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

func pingService(hs HTTPService) ServiceStatus {
	// TODO: make sure hs.URL is actually a valid URL
	res, err := http.Get(hs.URL)
	// expecting the HTTP client to err out if the service is down
	// TODO: check or specific errors e.g DNS not resolved
	if err != nil {
		// TODO: add TestURL to the log
		log.Errorf("[%s], is DOWN", hs.URL)
		return ServiceDown
	}

	if res.StatusCode == hs.StatusCode {
		log.Infof("[%s], is UP", hs.URL)
		return ServiceUP
	} else {
		log.Errorf("[%s], is FAILING", hs.URL)
		return ServiceFailing
	}
}

func MonitService(ctx context.Context, service HTTPService) {
	log.Infof("starting to monitor [%s] ...", service.URL)
	for {
		select {
		// TODO: declare a default interval in case service.Interval is not defined
		case <-time.After(service.Interval):
			// we already log the ServiceStatus so the returned status can be
			// ignored until we need to do sth more sophisticated here
			pingService(service)
		case <-ctx.Done():
			log.Warnf("ceasing to monitor: [%s] ...", service.URL)
		}
	}
}
