package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type CommandHandler struct {
	l *log.Logger
}

func NewCommandHandler(l *log.Logger) *CommandHandler {
	return &CommandHandler{l}
}

func (g *CommandHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Incoming Command")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "What are you tryna do bruh?", http.StatusBadRequest)
		return
	}

	// run the incoming command here and respond

	g.l.Printf("Executing: %s\n", d)
	rw.Write([]byte("awesome, good job\n"))
	return
}
