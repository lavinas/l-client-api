package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/lavinas/l-client-api/util/logger"
)

var (
	sg = make(chan os.Signal)
	r = mux.NewRouter()
)

// run is a local function that execute listener and prepare for grecefully shutdown
func run() {
	logger.Info("starting")
	// Run our server in a goroutine so that it doesn't block.
	s := &http.Server{
		Handler:      r,
		Addr:         ":8001",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal("listener error", err)
		}
	}()
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(sg, os.Interrupt)
	sig := <-sg
	logger.Info(fmt.Sprintf("shutting down with %s", sig))
	tc, cc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cc()
	s.Shutdown(tc)
	os.Exit(0)		
}
