package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/chowdhuryrahulc/oman-task/src/main/config"
)

func Start(appConfig *config.Config) {
	// Create a new server
	server, err := NewServer(appConfig)
	if err != nil {
		log.Errorf("server failed to start: %s", err)
		panic(err)
	}

	// Make the server to listen to the incoming requests
	if err = server.ListenAndServe(); err != nil {
		log.Errorf("Server failed to start: %s", err)
		panic(err)
	}
}
