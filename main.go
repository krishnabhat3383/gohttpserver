package main

import (
	"bhatji/gohttpserver/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "main-logger", log.LstdFlags)
	g := handlers.NewRootHandler(log.New(os.Stdout,
		"root", log.LstdFlags))
	f := handlers.NewCommandHandler(log.New(os.Stdout,
		"exec", log.LstdFlags))

	// this is the custom handler multiplexer
	// decides which handler to use for handling the request
	sMux := http.NewServeMux()
	sMux.Handle("/exec", f)
	sMux.Handle("/", g)

	// this is the custom server object
	mServer := &http.Server{
		Addr:         ":13000",
		Handler:      sMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := mServer.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// this is how you can receive os signals and get notified
	sigChannel := make(chan os.Signal) // basic go channel syntax
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	// this is how to use the channel to get the signal
	// this code basically blocks until channel is occupied
	term := <-sigChannel
	l.Println("yikes, we need to sleep ig xoxo", term)

	// this is how timeout is handled, before we want to terminate
	// the server. you can also shut it down directly, but
	// we want to wait for some time for all procs to exit
	// to do: check out contexts in go
	cont, _ := context.WithTimeout(context.Background(), 30*time.Second)
	mServer.Shutdown(cont)

}
