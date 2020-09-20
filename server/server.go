package server

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":9000", nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server ping")
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}
