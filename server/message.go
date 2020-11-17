package server

import (
	"encoding/json"
	"github.com/HaRshA10D/go-hello-server/logger"
	"net/http"
	"strconv"
)

var messages = map[int]Message{
	0: {Text: "This call may be recorded for training purposes."},
	1: {Text: "What’s crackin’?"},
	2: {Text: "GOOOOOD MORNING, VIETNAM!"},
	3: {Text: "Yo!"},
	4: {Text: "You know who this is."},
	5: {Text: "Ghostbusters, whatya want?"},
	6: {Text: "I'm Batman."},
}

type Message struct {
	Text string `json:"text"`
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	errorFunc := func() {
		logger.Info("No sequence received")
		json.NewEncoder(w).Encode(Message{Text: "Hell yo, pass the correct sequence"})
	}

	sequence, found := r.URL.Query()["sequence"]
	if !found {
		errorFunc()
		return
	}

	s, err := strconv.Atoi(sequence[0])
	if err != nil {
		errorFunc()
		return
	}

	logger.Info("For sequence %d, response %s", s, messages[s%7])
	json.NewEncoder(w).Encode(messages[s%7])
}


