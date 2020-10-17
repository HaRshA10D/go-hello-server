package server

import (
	"fmt"
	"github.com/HaRshA10D/go-hello-server/config"
	"github.com/HaRshA10D/go-hello-server/logger"
	"net/http"
)

// Start starts the http server
func Start() {
	config.Load()
	configureLogger()

	http.HandleFunc("/ping", pingHandler)
	port := fmt.Sprintf(":%d", config.Server().Port)
	http.ListenAndServe(port, nil)
}

func configureLogger() {
	logger.Init(*config.Logger())
	logger.Info("Starting server %d", config.Server().Port)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger().WithFields(map[string]interface{}{
		"path": "ping",
	})
	l.Info("Serving ping")
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}
