// Command localserve runs this shell's function.Main locally over plain
// HTTP, without the GCP Cloud Functions buildpack/Functions Framework.
// Run the prep-router action first to point internal/routersource at
// whichever router you want to serve - useful for local development and
// for CI that wants to exercise a real router without deploying anything.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	function "github.com/dash-xd/gospace-minimal/function"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		// Loopback only - this is meant to be reached from the same
		// machine/job that started it, not the network.
		host = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: function.Main,
	}

	serveErr := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
			return
		}
		close(serveErr)
	}()

	log.Printf("listening on http://%s", srv.Addr)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serveErr:
		log.Fatalf("serve: %v", err)

	case <-stop:
		log.Print("shutting down")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("shutdown: %v", err)
		}
	}
}
