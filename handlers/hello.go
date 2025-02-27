package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
} 

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("in /hello code")

	defer r.Body.Close()

	rawInput, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error occured reading input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", string(rawInput))
}