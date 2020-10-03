package server

import (
	"github.com/HaRshA10D/go-hello-server/config"
	"github.com/HaRshA10D/go-hello-server/logger"
	"net/http"
)

func Start() {
	config.Load()
	configureLogger()

	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":9000", nil)
}

func configureLogger() {
	logger.Init(*config.Logger())
	logger.Info("Starting server %d", 8080)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger().WithFields(map[string]interface{}{
		"path": "ping",
	})
	l.Info("Serving ping")
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}
