package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type RootHandler struct {
	l *log.Logger
}

func NewRootHandler(l *log.Logger) *RootHandler {
	return &RootHandler{l}
}

func (rt *RootHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rt.l.Println("guys, we've got a stray :)")

	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "bruh, bad request buddy, whaddya doin here?",
			http.StatusBadRequest)
		return
	}

	/// check for the right command here and
	// if it's right, return full status of the server
	rw.Write([]byte("not ready yet\n"))
}
