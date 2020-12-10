package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunHTTP(port int, router http.Handler) error {
	server := http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		Handler:      router,
		IdleTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	errs := make(chan error, 2)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errs <- err
			return
		}
		errs <- nil
	}()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		signal.Notify(quit, os.Kill)

		<-quit

		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(timeout); err != nil {
			errs <- err
			return
		}

		errs <- timeout.Err()
	}()

	return <-errs
}
