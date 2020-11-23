package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hjhussaini/url-shortener/logger"
)

func RunHTTP(address string, router http.Handler) {
	server := http.Server{
		Addr:         address,
		Handler:      router,
		IdleTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()
	logger.Info("Listening on", address)

	killSignal := make(chan os.Signal)
	signal.Notify(killSignal, os.Interrupt)
	signal.Notify(killSignal, os.Kill)

	signaled := <-killSignal
	logger.Info("Shutting server down (%sed)\n", signaled)

	// Gracefully shut down the server
	timeout, _ := context.WithTimeout(context.Background(), time.Second)
	server.Shutdown(timeout)
}
