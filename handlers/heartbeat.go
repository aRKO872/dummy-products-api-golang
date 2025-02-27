package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Heartbeat struct {
	l *log.Logger
}

func NewHeartbeat(l *log.Logger) *Heartbeat {
	return &Heartbeat{l}
}

func (h *Heartbeat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working Fine")
}