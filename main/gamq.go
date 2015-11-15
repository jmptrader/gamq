package main

import (
	"flag"

	"github.com/FireEater64/gamq"
	log "github.com/cihub/seelog"
)

func main() {
	// Set up a done channel, that's shared by the whole pipeline.
	// Closing this channel will kill all pipeline goroutines
	//done := make(chan struct{})
	//defer close(done)

	// Set up logging
	initializeLogging()

	// In the event of an unexpected shutdown - flush the log
	defer log.Flush()

	// Parse the command line flags
	config := parseCommandLineFlags()

	log.Info("Broker started.\n")

	connectionManager := gamq.ConnectionManager{}
	connectionManager.Initialize(&config)
}

func initializeLogging() {
	logger, err := log.LoggerFromConfigAsFile("config/logconfig.xml")

	if err != nil {
		log.Criticalf("An error occurred whilst initializing logging\n", err.Error())
		panic(err)
	}

	log.ReplaceLogger(logger)
}

func parseCommandLineFlags() gamq.Config {
	configToReturn := gamq.Config{}

	flag.IntVar(&configToReturn.Port, "port", 48879, "The port to listen on")

	flag.Parse()

	return configToReturn
}
