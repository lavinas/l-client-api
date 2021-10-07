package app

import (
	"net/http"
	"time"
	"os/signal"
	"context"
	"os"
	"log"
	"github.com/gorilla/mux"
)

var (
	l = log.New(os.Stdout, "product-api-", log.LstdFlags)
	sg = make(chan os.Signal)
	r = mux.NewRouter()
)

// run is a local function that execute listener and prepare for grecefully shutdown
func run() {
	l.Println("Starting app")
	// Run our server in a goroutine so that it doesn't block.
	s := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(sg, os.Interrupt)
	sig := <-sg
	tc, cc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cc()
	s.Shutdown(tc)
	l.Println("Shutting Down app", sig)
	os.Exit(0)		
}
