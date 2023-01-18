package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/logging"
	"github.com/karolsudol/vitalik/maker"
)

func main() {

	const projectID = "gcp-project-id"
	const subID = "payments-prod-sub"

	ctx := context.Background()

	// Creates a client.
	clientLogger, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer clientLogger.Close()

	// Sets the name of the log to write to.
	logName := subID

	loggerInfo := clientLogger.Logger(logName).StandardLogger(logging.Info)
	loggerErr := clientLogger.Logger(logName).StandardLogger(logging.Error)

	flag.Parse()

	maker := maker.Maker{
		SubID:     subID,
		ProjectID: projectID,
		LogInfo:   loggerInfo,
	}

	err = maker.PullMsgs()
	if err != nil {
		loggerErr.Println(err)
	}
}
