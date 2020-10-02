package server

import (
	"github.com/HaRshA10D/go-hello-server/logger"
	"net/http"
	"os"
)

func Start() {
	// Initialize logger
	loggerConfig := logger.Config{
		Level: logger.InfoLevel,
		Output: os.Stdout,
	}
	logger.Init(loggerConfig)
	logger.Info("Starting server %d", 8080)

	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":9000", nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger().WithFields(map[string]interface{}{
		"path": "ping",
	})
	l.Info("Serving ping")
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}
