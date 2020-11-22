package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/muaazsaleem/go-monit/monit"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// TODO: read Service Definitions from a file
	magnificentService := monit.HTTPService{
		URL:        "http://localhost:12345",
		StatusCode: 200,
		Interval:   time.Second * 20,
	}

	go handleSigterm(cancel)
	monit.MonitService(ctx, magnificentService)
}

// handleSigterm handles SIGTERM signal sent to the process.
func handleSigterm(cancelFunc func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	log.Info("Received Term signal. Terminating...")
	cancelFunc()
}
