package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunHTTP(address string, router http.Handler) {
	server := http.Server{
		Addr:         address,
		IdleTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()
	fmt.Println("Listening on", address)

	killSignal := make(chan os.Signal)
	signal.Notify(killSignal, os.Interrupt)
	signal.Notify(killSignal, os.Kill)

	signaled := <-killSignal
	fmt.Printf("Shutting server down (%s)", signaled)

	// Gracefully shut down the server
	timeout, _ := context.WithTimeout(context.Background(), time.Second)
	server.Shutdown(timeout)
}
